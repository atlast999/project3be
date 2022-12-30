package token

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JWTMaker struct {
	SecretKey string
}

func NewJWTMaker(secretKey string) (TokenMaker, error) {
	var err error
	return &JWTMaker{
		SecretKey: secretKey,
	}, err
}

func (maker *JWTMaker) CreateToken(userID int, username string, duration time.Duration) (string, error) {
	payload := NewPayload(userID, username, duration)
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token, err := jwtToken.SignedString([]byte(maker.SecretKey))
	return token, err
}

func (maker *JWTMaker) VerifyToken(token string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrTokenInvalid
		}
		return []byte(maker.SecretKey), nil
	}
	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		return nil, err
	}
	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, ErrTokenInvalid
	}
	return payload, nil
}
