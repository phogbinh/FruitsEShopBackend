package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var SqlDb *sql.DB

func Init() {
	var err error
	SqlDb, err = sql.Open("mysql", "root:freddy3601@tcp(127.0.0.1:3306)/user_db")
	if err != nil {
		log.Fatal(err.Error())
	}

	err = SqlDb.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
}
