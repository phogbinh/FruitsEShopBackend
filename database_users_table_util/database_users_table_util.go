package database_users_table_util

import "database/sql"

const (
	TableName              = "users"
	MailColumnName         = "Mail"
	PasswordColumnName     = "Password"
	UserNameColumnName     = "UserName"
	NicknameColumnName     = "Nickname"
	FnameColumnName        = "Fname"
	LnameColumnName        = "Lname"
	PhoneColumnName        = "Phone"
	LocationColumnName     = "Location"
	MoneyColumnName        = "Money"
	IntroductionColumnName = "Introduction"
)

// CreateDatabaseUsersTableIfNotExists creates table `users`.
func CreateDatabaseUsersTableIfNotExists(databasePtr *sql.DB) error {
	_, createTableError := databasePtr.Exec("CREATE TABLE IF NOT EXISTS user_data.users" +
		"(" +
		MailColumnName + "			VARCHAR(320)	NOT NULL," +
		PasswordColumnName + "		VARCHAR(30)		NOT NULL," +
		UserNameColumnName + "		VARCHAR(30)		NOT NULL," +
		NicknameColumnName + "		VARCHAR(30)				," +
		FnameColumnName + "			VARCHAR(15)		NOT NULL," +
		LnameColumnName + "			VARCHAR(15)		NOT NULL," +
		PhoneColumnName + "			VARCHAR(30)				," +
		LocationColumnName + "		VARCHAR(30)				," +
		MoneyColumnName + "			DECIMAL(30, 2)			," +
		IntroductionColumnName + "	VARCHAR(255)			," +
		"PRIMARY KEY(" + UserNameColumnName + ")," +
		"UNIQUE(" + MailColumnName + ")," +
		"UNIQUE(" + NicknameColumnName + ")" +
		");")
	return createTableError
}
