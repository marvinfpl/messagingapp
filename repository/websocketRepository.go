package repository

import (
	"context"
	"encoding/json"
	"errors"
	"messagingapp/commons"
	"messagingapp/models"
	"sync"
	"github.com/redis/go-redis/v9"
	"github.com/gorilla/websocket"
)

type WebsocketRepository struct {
	Conn		*websocket.Conn
	Map		 	sync.Map
	Wg          sync.WaitGroup
	Rds			*redis.Client
	Broadcast	chan models.Message
	MsgDB   	MessageRepositoryDB
}

type WebsocketService interface {
}

func NewWebsocketRepository() *WebsocketRepository {
	return &WebsocketRepository{
		Conn: 		nil,
		Map:		sync.Map{},
		Wg:			sync.WaitGroup{},
		Rds: 		commons.InitRedis(),
		MsgDB:		MessageRepositoryDB{
			Coll: commons.InitMongo("message"),
		},
	}
}

func (ws *WebsocketRepository) ReadMessage() {
	for {
		var msg models.Message
		err := ws.Conn.ReadJSON(&msg)
		if err != nil {
			break
		}
		ws.Broadcast <- msg
		ws.Wg.Add(1)
	}
}

func (ws *WebsocketRepository) WriteMessage(uuid string) error {
	for {
		ws.Wg.Wait()
		msg := <- ws.Broadcast
		receiver, ok := ws.Map.Load(msg.To)
		if ok {
			messages, err := ws.ReadUnsentMessage(uuid)
			if err != nil {
				break
			}
			for _, msg := range messages {
				err := receiver.(*websocket.Conn).WriteJSON(msg) // send message
				if err != nil {
					break
				}
				_, err = ws.MsgDB.Coll.InsertOne(context.TODO(), msg) // register message
				if err != nil {
					break
				}
			}
		} else {
			errChan := make(chan error)
			go func() {
				err := ws.WriteUnsentMessage(uuid, msg) // use redis for deleting old record instead of KAFFEKA
				if err != nil {
					errChan <- err
					close(errChan)
				}
			}()
			err := <- errChan
			if err != nil {
				break
			}
		}
		ws.Wg.Done()
	}
	return nil
}

func (ws *WebsocketRepository) ReadUnsentMessage(uuid string) ([]models.Message, error) {
	var messages []models.Message

	for {
		result, err := ws.Rds.RPop(context.Background(), uuid).Result()
		if err == redis.Nil {
			break
		} else if err != nil {
			return nil, err
		}
		var msg models.Message
		err = json.Unmarshal([]byte(result), &msg)
		if err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}
	return messages, nil
}

func (ws *WebsocketRepository) WriteUnsentMessage(uuid string, msg models.Message) error {
	messageBytes, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	err = ws.Rds.LPush(context.Background(), uuid, messageBytes).Err()
	if err != nil {
		return errors.New("cant store messages in redis: " + err.Error())
	}

	return nil
}