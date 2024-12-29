package commons

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"context"
)

func InitMongo(collection string) *mongo.Collection {
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URI"))
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic("cannot connect to mongo: " + err.Error())
	}
	return client.Database("messagingapp.db").Collection(collection)
}