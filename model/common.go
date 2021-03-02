package model

import (
	"time"
)

type Model struct {
	ID        string `gorm:"primaryKey"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

type SrvError struct {
	Error string `json:"error"`
}
