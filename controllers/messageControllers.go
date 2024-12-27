package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/gorilla/websocket"
	"net/http"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func SendMessage(c echo.Context) error {
	_, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "cannot upgrade connection"+err.Error()})
	}
	return nil
}