package repository

import (
	"messagingapp/commons"
	"messagingapp/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"context"
	"time"
)

type MessageRepositoryDB struct {
	Coll *mongo.Collection
}

type MessageService interface {
	CreateMessage(models.Message) error
	DeleteMessage(models.Message) error
	UpdateMessage(models.Message) error
}

func NewMessageRepositoryDB() *MessageRepositoryDB {
	return &MessageRepositoryDB{
		Coll: commons.InitMongo("message"),
	}
}

func (r *MessageRepositoryDB) CreateMessage(message *models.Message) error {
	message.CreatedAt = time.Now()
	_, err := r.Coll.InsertOne(context.TODO(), message)
	return err
}

func (r *MessageRepositoryDB) DeleteMessage(id uint) error {
	filter := bson.M{"ID": id}
	_, err := r.Coll.DeleteOne(context.TODO(), filter)
	return err
}

func (r *MessageRepositoryDB) UpdateMessage(message models.Message) error {
	_, err := r.Coll.UpdateByID(context.TODO(), message.ID, message.Content)
	return err
}

func (r *MessageRepositoryDB) GetMessages(chat models.Chat) ([]models.Message, error) {
	filter := bson.M{"UUID": chat.UUID}
	cur, err := r.Coll.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())

	var messages []models.Message
	for cur.Next(context.TODO()) {
		var msg models.Message
		err := cur.Decode(msg)
		if err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}

	return messages, nil
}