package model

import (
	"time"
)

type UserModel struct {
	ID          uint64    `json:"id" validate:"required"`
	Name        string    `json:"name" validate:"required"`
	MailAddress string    `json:"mail" validate:"required,email"`
	Birthday    time.Time `json:"birthday" validate:"datetime"`
}
