package database

import (
	"context"
	"go_api_boilerplate/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

const (
	CtxKeyDB string = "__database__"
)

func NewMysqlHandler() DBHandler {
	c := config.NewConfig()
	dc := &DBConfig{
		Host:     c.DB.Host,
		Username: c.DB.Username,
		Password: c.DB.Password,
		DBName:   c.DB.DBName,
		Port:     c.DB.Port,
	}
	return &mysqlHandlerImpl{dbConfig: dc}
}

type mysqlHandlerImpl struct {
	dbConfig *DBConfig
}

func (m *mysqlHandlerImpl) Open(ctx context.Context) (*gorm.DB, error) {
	if db, ok := ctx.Value(CtxKeyDB).(*gorm.DB); ok {
		return db, nil
	}
	return m.NewDB(ctx)
}

func (m *mysqlHandlerImpl) Close(db *gorm.DB) {
	sqlDB, _ := db.DB()
	if err := sqlDB.Close(); err != nil {
		panic(err)
	}
	log.Println("db close")
}

func (m *mysqlHandlerImpl) mysqlDsn() string {
	dsn := m.dbConfig.Username + ":" + m.dbConfig.Password + "@tcp(" + m.dbConfig.Host + ":" + m.dbConfig.Port + ")/" + m.dbConfig.DBName + "?charset=utf8mb4 &parseTime=True&loc=Local"
	return dsn
}

func (m *mysqlHandlerImpl) config() *gorm.Config {
	return &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}
}

func (m *mysqlHandlerImpl) NewDB(ctx context.Context) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(m.mysqlDsn()), m.config())
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(50)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)
	return db, err
}
