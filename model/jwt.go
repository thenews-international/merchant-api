package model

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

const (
	CtxKeyXUser ctxKey = "x-user"
)

type ctxKey string

type CtxUser struct {
	UserId         uuid.UUID
	TokenCreatedAt *time.Time
}

type Token struct {
	UserId    string     `json:"userId"`
	Email     string     `json:"email"`
	CreatedAt *time.Time `json:"createdAt"`
	*jwt.StandardClaims
}

// CtxUser should not be returned as a pointer
// due to it's unable to read via context.WithValue
func (t *Token) ToCtxUser() CtxUser {
	id, _ := uuid.Parse(t.UserId)
	return CtxUser{
		UserId:         id,
		TokenCreatedAt: t.CreatedAt,
	}
}
