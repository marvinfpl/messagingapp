package controllers

import (
	_"github.com/labstack/echo/v4"
	_"go.mongodb.org/mongo-driver/mongo"
	_"go.mongodb.org/mongo-driver/bson"
	_"messagingapp/models"
	_"messagingapp/repository"
)

/**var (
	ch = repository.NewChatRepositoryDB()
)

func GetChat(c echo.Context) error {
	uuid := c.Param("uuid")
	ch.GetChat(uuid)
}**/ // still html need