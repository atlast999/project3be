package token

import (
	"time"

	"github.com/o1egl/paseto"
)

type PasetoMaker struct {
	SecretKey string
	Paseto paseto.V2
}

func NewPasetoMaker(secretKey string) (TokenMaker, error) {
	var err error
	return &PasetoMaker{
		SecretKey: secretKey,
		Paseto: *paseto.NewV2(),
	}, err
}

func (maker *PasetoMaker) CreateToken(userID int, username string, duration time.Duration) (string, error) {
	payload := NewPayload(userID, username, duration)
	return maker.Paseto.Encrypt([]byte(maker.SecretKey), payload, nil)
}

func (maker *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}
	err := maker.Paseto.Decrypt(token, []byte(maker.SecretKey), payload, nil)
	if err != nil {
		return nil, ErrTokenInvalid
	}
	err = payload.Valid()
	if err != nil {
		return nil, err
	}
	return payload, nil
}
