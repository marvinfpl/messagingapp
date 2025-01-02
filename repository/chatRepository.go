package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"context"
	"messagingapp/commons"
	"messagingapp/models"
	"errors"
)

type ChatRepositoryDB struct {
	Coll *mongo.Collection
}

type ChatService interface {
	CreateChat(*models.Chat) error
	DeleteChat(string) error
	AddUsers(*models.Chat, []uint) error
	RemoveUsers(*models.Chat, []uint) error
}

func NewChatRepositoryDB() *ChatRepositoryDB {
	return &ChatRepositoryDB{
		Coll: commons.InitMongo("chat"),
	}
}

func (r *ChatRepositoryDB) CreateChat(chat *models.Chat) error {
	chat.UUID = commons.GenerateUUID()
	_, err := r.Coll.InsertOne(context.TODO(), chat)
	return err
}

func (r *ChatRepositoryDB) DeleteChat(uuid string) error {
	filter := bson.M{"UUID": uuid}
	_, err := r.Coll.DeleteOne(context.TODO(), filter)
	return err
}

func (r *ChatRepositoryDB) GetChat(uuid string) (models.Chat, error) {
	filter := bson.M{"UUID": uuid}
	result := r.Coll.FindOne(context.TODO(), filter)
	var chat models.Chat
	err := result.Decode(&chat)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.Chat{}, errors.New("chat doesn't exists")
		}
		return models.Chat{}, err
	}
	return chat, nil
}

/**func (r *ChatRepositoryDB) AddUsers(chat *models.Chat, userIDs []uint) error {
	var IDs []interface{}
	for _, id := range userIDs {
		if !commons.Contains(chat.UserIDs, id) {
			IDs = append(IDs, id)
		}
	}

	filter := bson.M{"UUID": chat.UUID}
	_, err := r.coll.UpdateOne(context.TODO(), filter, IDs)
	return err
}

func (r *ChatRepositoryDB) RemoveUsers(chat *models.Chat, userIDs []uint) error {
	var IDs []interface{}
	for _, id := range userIDs {
		if commons.Contains(chat.UserIDs, id) {
			IDs = append(IDs, id)
		}
	}

	filter := bson.M{"UUID": chat.UUID}
	_, err := r.coll.UpdateOne(context.TODO(), filter, IDs)
	return err
}**/ // currently out of source cuz don't need that for naw