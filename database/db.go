package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

var db *gorm.DB
var (
	host   = ""
	port   = ""
	user   = ""
	pwd    = ""
	dbName = ""
)

func getLogger() logger.Interface {
	return logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			logLevel: logger.Info
		})
}

func Connect() (err error) {
	db := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, pwd, dbName)
	db, err :=
}

func Get() (err error) {

}
