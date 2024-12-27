package data

import (
	"go.mongodb.org/mongo-driver/mongo"
	_"go.mongodb.org/mongo-driver/bson"
	m"messagingapp/models"
	"context"
)

var (
	ctx = context.Background()
)

type ChatRep struct {
	db *mongo.Collection
}

func (r *ChatRep) Get(uuid string) m.ChatRoom {
	
}