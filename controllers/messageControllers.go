package controllers

import (
	_"messagingapp/models"
	"messagingapp/repository"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	ws = repository.NewWebsocketRepository()
	//m = repository.NewMessageRepositoryDB()
	//ch = repository.NewChatRepositoryDB()
)

func SendMessage(c echo.Context) error {
	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "cannot upgrade connection: " + err.Error()})
	}
	ws.Conn = conn

	userName := c.Get("user_id").(uint)
	uuid := c.Param("uuid")
	ws.Map.Store(userName, conn)
	defer ws.Map.Delete(userName)
	defer conn.Close()

	go ws.ReadMessage() // use REDIS
	go ws.WriteMessage(uuid)

	return c.JSON(http.StatusOK, echo.Map{})
}

/**func GetMessages(c echo.Context) error {
	// i think it maybe some html that gets messages in input and do some nginx shits in ui
	uuid := c.Param("uuid")
	chat, err := ch.GetChat(uuid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "chat doesn't exists: " + err.Error()})
	}

	m.GetMessages(chat)

}**/ // html version of messagingapp