package controllers

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var JwtKey = []byte(os.Getenv("Secret"))

type Claim struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func GenerateToken(email string) (map[string]string, error) {
	expTime := time.Now().Add(1 * time.Hour)
	claim := &Claim{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		return nil, err
	}
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rfclaims := refreshToken.Claims.(jwt.MapClaims)
	rfclaims["email"] = email
	rfclaims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	refresh, err := refreshToken.SignedString([]byte("secret"))
	if err != nil {
		return nil, err
	}
	return map[string]string{
		"Token":    tokenString,
		"referesh": refresh,
	}, nil
}

var Val string

func Validate(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&Claim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(JwtKey), nil
		},
	)
	if err != nil {
		return
	}
	claim, ok := token.Claims.(*Claim)
	Val = claim.Email
	if !ok {
		err = errors.New("could not parse claims")
		return
	}
	if claim.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}
	return
}
