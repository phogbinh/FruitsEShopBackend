package main

import (
	"database/sql"
	"log"
	"os"

	db "backend/database"
	discountPoliciesTable "backend/database_discount_policies_tables_util/database_discount_policies_table_util"
	discountPolicyTypesTable "backend/database_discount_policies_tables_util/database_discount_policy_types_table_util"
	seasoningsDiscountPoliciesTable "backend/database_discount_policies_tables_util/database_seasonings_discount_policies_table_util"
	shippingDiscountPoliciesTable "backend/database_discount_policies_tables_util/database_shipping_discount_policies_table_util"
	specialEventDiscountPoliciesTable "backend/database_discount_policies_tables_util/database_special_event_discount_policies_table_util"
	DUTU "backend/database_users_table_util"
	"backend/router"

	"github.com/gin-gonic/gin"
)

const createTableErrorMessage = "Error creating database table: %q."

func getDatabaseHandler() *sql.DB {
	databasePtr, openDatabaseError := sql.Open("mysql", os.Getenv("DATABASE_URL"))
	if openDatabaseError != nil {
		log.Fatalf("Error opening database: %q.", openDatabaseError)
	}
	return databasePtr
}

func run() {
	databasePtr := getDatabaseHandler()
	defer databasePtr.Close()
	createDatabaseUsersTableError := DUTU.CreateTableIfNotExists(databasePtr)
	if createDatabaseUsersTableError != nil {
		log.Fatalf(createTableErrorMessage, createDatabaseUsersTableError)
	}
	initializeError := discountPolicyTypesTable.Initialize(databasePtr)
	if initializeError != nil {
		log.Fatalf("Error initializing database discount policy types table: %q.", initializeError)
	}
	var createTableError error
	createTableError = discountPoliciesTable.CreateTableIfNotExists(databasePtr)
	if createTableError != nil {
		log.Fatalf(createTableErrorMessage, createTableError)
	}
	createTableError = shippingDiscountPoliciesTable.CreateTableIfNotExists(databasePtr)
	if createTableError != nil {
		log.Fatalf(createTableErrorMessage, createTableError)
	}
	createTableError = seasoningsDiscountPoliciesTable.CreateTableIfNotExists(databasePtr)
	if createTableError != nil {
		log.Fatalf(createTableErrorMessage, createTableError)
	}
	createTableError = specialEventDiscountPoliciesTable.CreateTableIfNotExists(databasePtr)
	if createTableError != nil {
		log.Fatalf(createTableErrorMessage, createTableError)
	}
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
