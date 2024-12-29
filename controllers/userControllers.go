package controllers

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"messagingapp/commons"
	"messagingapp/models"
	"messagingapp/repository"
)

var (
	r = repository.NewUserRepositoryDB()
)


func Register(c echo.Context) error {
	user := new(models.User)
	c.Bind(user) // verifier si les primarykey qui sont respectés: if user.Email in db -> signin -> a mettre au testingcase
	if err := c.Validate(user); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "invalid data"+err.Error()})
	}

	hash, err := commons.HashPassword(user.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "cannot hash the password"})
	}
	user.Password = hash

	err = r.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "cannot create user"})
	}

	token, err := commons.GenerateJwt(user.Name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "cannot generate jwt"})
	}

	return c.JSON(http.StatusCreated, echo.Map{"message": "user successfully registered", "user": user, "token": token})

}

func Login(c echo.Context) error {
	userForm := new(models.User)
	c.Bind(&userForm)
	err := c.Validate(&userForm)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "invalid data"+err.Error()})
	}

	user, err := r.GetUser(userForm.Email)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "user not found"})
	}
	ok := commons.ComparePassword(user.Password, userForm.Password)
	if !ok {
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": "incorrect password"})
	}

	token, err := commons.GenerateJwt(user.Name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "cannot generate jwt"})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "user successfully logged in", "token": token})
}

// func UpdateUser(c echo.Context) error {}