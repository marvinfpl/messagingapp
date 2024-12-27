package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
	c"messagingapp/controllers"
)

func main() {
	e := echo.New()

	middleware.Logger()
	middleware.Recover()

	//e.GET("/:uuid", )
	e.POST("/register", c.Register)
	e.POST("/login", c.Login)

	e.Logger.Fatal(e.Start(os.Getenv("PORT")))
}