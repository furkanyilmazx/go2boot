package system

import (
	"database/sql"
	"fmt"
	"go2api/system/config/models"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB(dbConfig *models.DB) *sql.DB {
	dbConf := fmt.Sprintf("%s:%s@%s/%s", dbConfig.Username, dbConfig.Password,dbConfig.Url, dbConfig.Name)
	db, err := sql.Open(dbConfig.Driver, dbConf) // this does not really open a new connection

	if err != nil {
		log.Fatalf("Error on initializing database connection: %s", err.Error())
	}

	err = db.Ping() // This DOES open a connection if necessary. This makes sure the database is accessible
	if err != nil {
		log.Fatalf("Error on opening database connection: %s", err.Error())
	}
	return db
}
