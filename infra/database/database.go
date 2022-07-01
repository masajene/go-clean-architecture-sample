package database

import (
	"context"
	"gorm.io/gorm"
)

// DBConfig struct
type DBConfig struct {
	Host     string
	Username string
	Password string
	DBName   string
	Port     string
}

type DBHandler interface {
	Open(ctx context.Context) (*gorm.DB, error)
	Close(db *gorm.DB)
}
