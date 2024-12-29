package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"context"
	"messagingapp/commons"
	"messagingapp/models"
)

type ChatRepositoryDB struct {
	coll *mongo.Collection
}

type ChatService interface {
	CreateChat(*models.Chat) error
	DeleteChat(string) error
	AddUsers(*models.Chat, []uint) error
	RemoveUsers(*models.Chat, []uint) error
}

func NewChatRepositoryDB() *ChatRepositoryDB {
	return &ChatRepositoryDB{
		coll: commons.InitMongo("chat"),
	}
}

func (r *ChatRepositoryDB) CreateChat(chat *models.Chat) error {
	chat.UUID = commons.GenerateUUID()
	_, err := r.coll.InsertOne(context.TODO(), chat)
	return err
}

func (r *ChatRepositoryDB) DeleteChat(uuid string) error {
	filter := bson.M{"UUID": uuid}
	_, err := r.coll.DeleteOne(context.TODO(), filter)
	return err
}

func (r *ChatRepositoryDB) AddUsers(chat *models.Chat, userIDs []uint) error {
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
}