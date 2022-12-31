package token

import (
	"errors"
	"time"
)

var (
	ErrTokenExpired = errors.New("Token has expired")
	ErrTokenInvalid = errors.New("Token is invalid")
)

type Payload struct {
	UserID    int32       `json:"id"`
	Username  string    `json:"username"`
	ExpiredAt time.Time `json:"expired_at"`
}

func (payload *Payload) Valid() error {
	if payload.ExpiredAt.Before(time.Now()) {
		return ErrTokenExpired
	}
	return nil
}

func NewPayload(userId int32, username string, duration time.Duration) *Payload {
	return &Payload{
		UserID:    userId,
		Username:  username,
		ExpiredAt: time.Now().Add(duration),
	}
}

type TokenMaker interface {
	CreateToken(userID int32, username string, duration time.Duration) (string, error)
	VerifyToken(token string) (*Payload, error)
}
