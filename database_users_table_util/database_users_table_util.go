package database_users_table_util

import "database/sql"

const (
	TableName          = "users"
	UserNameColumnName = "username"
	PasswordColumnName = "password"
)

// CreateDatabaseUsersTableIfNotExists creates a table named 'users' for the given database pointer if the table has not already existed.
func CreateDatabaseUsersTableIfNotExists(databasePtr *sql.DB) error {
	_, createTableError := databasePtr.Exec("CREATE TABLE IF NOT EXISTS users (" + UserNameColumnName + " VARCHAR(255) PRIMARY KEY, " + PasswordColumnName + " VARCHAR(255) NOT NULL)")
	return createTableError
}
