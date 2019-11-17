package main

import (
	db "backend/database"
	"backend/router"
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
)

func getDatabaseHandler() *sql.DB {
	databasePtr, openDatabaseError := sql.Open("mysql", "root:Thucchinh1312@@tcp(127.0.0.1:3306)/user_data")
	if openDatabaseError != nil {
		log.Fatalf("Error opening database: %q.", openDatabaseError)
	}
	return databasePtr
}

func run() {
	databasePtr := getDatabaseHandler()
	defer databasePtr.Close()
	db.SqlDb = databasePtr
	db.CreateDatabases(db.SqlDb)

	var httpServer *gin.Engine

	httpServer = gin.Default()

	router.Register(httpServer, databasePtr)

	serverAddr := "0.0.0.0:8080"

	// listen and serve on 0.0.0.0:8080
	httpServer.Run(serverAddr)
}

func main() {
	run()
}
