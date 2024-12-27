package controllers

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"messagingapp/data"
)

var (
	cr = data.ChatRep{}
)

func ChatRoom(c echo.Context) error {
	uuid := c.Param("uuid")
	filter := bson.M{"UUID": uuid}
	cr.Find()

}