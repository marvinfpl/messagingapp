package commons

import (
	"github.com/golang-jwt/jwt/v5"
	"os"
)

// custom claims ?

var (
	secretKey = os.Getenv("SECRETKEY")
)

func GenerateJwt(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return t, nil
}

//func middlewareJwt(c echo.Context)