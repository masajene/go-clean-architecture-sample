package infrastructure

import (
	"os"

	"github.com/jinzhu/gorm"
)

// SQLHandler sql handler struct
type SQLHandler struct {
	Db *gorm.DB
}

// CreateHandler return sql handler
func CreateHandler() *SQLHandler {
	filePath := os.Getenv("SQL_FILE_PATH")
	db, err := gorm.Open("sqlite3", filePath)
	if err != nil {
		panic(err.Error())
	}

	sqlHandler := new(SQLHandler)
	sqlHandler.Db = db
	return sqlHandler
}
