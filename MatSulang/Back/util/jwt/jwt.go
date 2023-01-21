package util

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type AuthTokenClaims struct {
	UserID   string `json:"userID"`
	Nickname string `json:"NickName"`
	jwt.StandardClaims
}

func CreateToken(userID string) (string, error) {

	os.Setenv("ACCESS_SECRET", "matsulang")
	atClaims := jwt.MapClaims{}

	atClaims["authorized"] = true
	atClaims["user_id"] = userID
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodES256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}
