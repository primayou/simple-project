package handler

import (
	"crypto/sha1"
	"encoding/hex"
	"simple-projects/models"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func encrypt(pass string) string {
	h := sha1.New()
	h.Write([]byte(pass))
	return hex.EncodeToString(h.Sum(nil))
}

func CreateAccessToken(user *models.User) (string, error) {

	not_before_date := time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix()
	sign := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid": user.ID,
		"nbf": not_before_date,
	})

	// Sign and get the complete encoded token as a string using the secret
	token, err := sign.SignedString([]byte("secret"))

	if err != nil {
		return "", err
	}

	return token, nil
}
