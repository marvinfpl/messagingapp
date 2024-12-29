package repository

import (
	"messagingapp/commons"
	"messagingapp/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"context"
)

type MessageRepositoryDB struct {
	coll *mongo.Collection
}

type MessageService interface {
	CreateMessage(*models.Message) error
	DeleteMessage(*models.Message) error
	UpdateMessage(*models.Message) error
}

func NewMessageRepositoryDB() *MessageRepositoryDB {
	return &MessageRepositoryDB{
		coll: commons.InitMongo("message"),
	}
}

func (r *MessageRepositoryDB) CreateMessage(message *models.Message) error {
	_, err := r.coll.InsertOne(context.TODO(), message)
	return err
}

func (r *MessageRepositoryDB) DeleteMessage(id uint) error {
	filter := bson.M{"ID": id}
	_, err := r.coll.DeleteOne(context.TODO(), filter)
	return err
}

//func (r *MessageRepositoryDB) GetMessages(uuid string) error | dilemma | to put or not to put, that is the question

func (r *MessageRepositoryDB) UpdateMessage(message *models.Message) error {
	_, err := r.coll.UpdateByID(context.TODO(), message.ID, message.Content)
	return err
}