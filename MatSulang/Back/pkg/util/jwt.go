package util

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	UserID   string `json:"userID"`
	Nickname string `json:"NickName"`
	jwt.StandardClaims
}

func CreateToken(userID string) (string, error) {

	os.Setenv("ACCESS_SECRET", "matsulang")

	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		},
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		fmt.Printf("token err : %v", err)
		return "", err
	}
	return token, nil
}
