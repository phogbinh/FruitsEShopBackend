package database_users_table_util

import (
	"database/sql"
	"errors"
	"strconv"

	. "backend/model"
	"backend/util"
)

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

const (
	queryInsertUser         = "INSERT INTO " + TableName + " VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	queryGetAllUsers        = "SELECT * FROM " + TableName
	queryGetUserByUserName  = "SELECT * FROM " + TableName + " WHERE " + UserNameColumnName + " = ?"
	queryGetUserByMail      = "SELECT * FROM " + TableName + " WHERE " + MailColumnName + " = ?"
	queryUpdateUserPassword = "UPDATE " + TableName + " SET " + PasswordColumnName + " = ? WHERE " + UserNameColumnName + " = ?"
	queryDeleteUser         = "DELETE FROM " + TableName + " WHERE " + UserNameColumnName + " = ?"
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
		LocationColumnName + "		VARCHAR(255)			," +
		MoneyColumnName + "			DECIMAL(30, 2)			," +
		IntroductionColumnName + "	VARCHAR(255)			," +
		"PRIMARY KEY(" + UserNameColumnName + ")," +
		"UNIQUE(" + MailColumnName + ")," +
		"UNIQUE(" + NicknameColumnName + ")" +
		");")
	return createTableError
}

// InsertUser inserts the given user to the database `users` table.
func InsertUser(user User, databasePtr *sql.DB) Status {
	return queryDatabase(databasePtr, queryInsertUser, user.Mail, user.Password, user.UserName, user.Nickname, user.Fname, user.Lname, user.Phone, user.Location, user.Money, user.Introduction)
}

// GetAllUsers returns all users' information.
func GetAllUsers(databasePtr *sql.DB) ([]User, Status) {
	queryRowsPtr, queryError := databasePtr.Query(queryGetAllUsers)
	if queryError != nil {
		return nil, util.StatusInternalServerError(GetAllUsers, queryError)
	}
	defer queryRowsPtr.Close()
	return getAllUsers(queryRowsPtr)
}

func getAllUsers(databaseUsersTableRowsPtr *sql.Rows) ([]User, Status) {
	var users []User
	for databaseUsersTableRowsPtr.Next() {
		var user User
		scanError := databaseUsersTableRowsPtr.Scan(&user.Mail, &user.Password, &user.UserName, &user.Nickname, &user.Fname, &user.Lname, &user.Phone, &user.Location, &user.Money, &user.Introduction)
		if scanError != nil {
			return nil, util.StatusInternalServerError(getAllUsers, scanError)
		}
		users = append(users, user)
	}
	return users, util.StatusOK()
}

func getUserByKeyColumnFromDatabaseUsersTable(queryGetUserByKeyColumn string, keyColumnValue string, databasePtr *sql.DB) (User, Status) {
	var dumpUser User
	users, getStatus := getUsersByKeyColumnFromDatabaseUsersTable(queryGetUserByKeyColumn, keyColumnValue, databasePtr)
	if !util.IsStatusOK(getStatus) {
		return dumpUser, getStatus
	}
	if len(users) != 1 {
		return dumpUser, util.StatusInternalServerError(getUserByKeyColumnFromDatabaseUsersTable, errors.New("Query 1 user but got "+strconv.Itoa(len(users))+" user(s) instead."))
	}
	return users[0], util.StatusOK()
}

func getUsersByKeyColumnFromDatabaseUsersTable(queryGetUserByKeyColumn string, keyColumnValue string, databasePtr *sql.DB) ([]User, Status) {
	queryRowsPtr, queryError := databasePtr.Query(queryGetUserByKeyColumn, keyColumnValue)
	if queryError != nil {
		return nil, util.StatusInternalServerError(getUsersByKeyColumnFromDatabaseUsersTable, queryError)
	}
	defer queryRowsPtr.Close()
	users, getStatus := getAllUsers(queryRowsPtr)
	if !util.IsStatusOK(getStatus) {
		return nil, getStatus
	}
	return users, util.StatusOK()
}

// UpdateUserPassword updates the given user's password to the database `users` table.
func UpdateUserPassword(userName string, userNewPassword string, databasePtr *sql.DB) Status {
	return queryDatabase(databasePtr, queryUpdateUserPassword, userNewPassword, userName)
}

// DeleteUser deletes the given user from the database `users` table.
func DeleteUser(userName string, databasePtr *sql.DB) Status {
	return queryDatabase(databasePtr, queryDeleteUser, userName)
}

// IsExistingUser determines whether the given user is in the database `users` table.
func IsExistingUser(userName string, databasePtr *sql.DB) (bool, Status) {
	users, getStatus := getUsersByKeyColumnFromDatabaseUsersTable(queryGetUserByUserName, userName, databasePtr)
	if !util.IsStatusOK(getStatus) {
		return false, getStatus
	}
	return len(users) > 0, util.StatusOK()
}

func queryDatabase(databasePtr *sql.DB, query string, executeArguments ...interface{}) Status {
	prepareStatementPtr, prepareError := databasePtr.Prepare(query)
	if prepareError != nil {
		return util.StatusInternalServerError(queryDatabase, prepareError)
	}
	_, executeError := prepareStatementPtr.Exec(executeArguments...)
	if executeError != nil {
		return util.StatusInternalServerError(queryDatabase, executeError)
	}
	return util.StatusOK()
}

// GetUserByMail returns an user's information by the given mail.
func GetUserByMail(mail string, databasePtr *sql.DB) (User, Status) {
	return getUserByKeyColumnFromDatabaseUsersTable(queryGetUserByMail, mail, databasePtr)
}

// GetUserByUserName returns an user's information by the given user name.
func GetUserByUserName(userName string, databasePtr *sql.DB) (User, Status) {
	return getUserByKeyColumnFromDatabaseUsersTable(queryGetUserByUserName, userName, databasePtr)
}
