package data

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	m "messagingapp/models"
)

type MessageRep struct {
	db *mongo.Collection
}

func (r *MessageRep) Create(message *m.Message) error {
	_, err := r.db.InsertOne(ctx, message)
	return err
}

func (r *MessageRep) Update(message *m.Message) error {
	_, err := r.db.ReplaceOne(ctx, bson.M{"_id": message.ID}, bson.M{"$set": bson.M{
		"content": message.Content,
		"from": message.From,
		"to": message.To,
		"created_at": message.CreatedAt,
	}})
	return err
}

func (r *MessageRep) Delete(message *m.Message) error {
	_, err := r.db.DeleteOne(ctx, bson.M{"_id": message.ID})
	return err
}