package config

import (
	"log"
	"os"
)

type Config struct {
	DB struct {
		Host     string
		Username string
		Password string
		DBName   string
		Port     string
	}
	Token string
}

// NewConfig create
func NewConfig() *Config {
	goenv := os.Getenv("GO_ENV_TYPE")
	log.Println("goenv:", goenv)
	// switch goenv {
	// case "production":
	// 	return newProdConfig()
	// case "develop":
	// 	return newDevConfig()
	// default:
	// 	return newProdConfig()
	// }
	return newLocalConfig()
}

// NewConfig config initialize
func newServerConfig() *Config {
	c := new(Config)
	c.DB.Host = "127.0.0.1"
	c.DB.Username = "root"
	c.DB.Password = "password"
	c.DB.DBName = "sales_manage_dev"
	c.DB.Port = "55001"
	c.Token = "l66gdz7rnp34"
	return c
}

// NewConfig config initialize
func newLocalConfig() *Config {
	c := new(Config)
	c.DB.Host = "127.0.0.1"
	c.DB.Username = "root"
	c.DB.Password = "password"
	c.DB.DBName = "sales_manage_dev"
	c.DB.Port = "55001"
	c.Token = "l66gdz7rnp34"
	return c
}
