package entity

import "time"

type User struct {
	ID          uint64    `json:"id" gorm:"column:id"`
	Name        string    `json:"name" gorm:"column:name"`
	MailAddress string    `json:"mail" gorm:"column:mail"`
	Birthday    time.Time `json:"birthday" gorm:"column:birthday"`
	CreatedAt   time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (m *User) TableName() string {
	return "users"
}
