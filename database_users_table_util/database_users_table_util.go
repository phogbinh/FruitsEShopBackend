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
	StaffFlagColumnName    = "StaffFlag"
)

const (
	queryInsertUser          = "INSERT INTO " + TableName + " VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	queryGetAllUsers         = "SELECT * FROM " + TableName
	queryGetUserByUserName   = "SELECT * FROM " + TableName + " WHERE " + UserNameColumnName + " = ?"
	queryGetUserByMail       = "SELECT * FROM " + TableName + " WHERE " + MailColumnName + " = ?"
	queryUpdateUserPassword  = "UPDATE " + TableName + " SET " + PasswordColumnName + " = ? WHERE " + UserNameColumnName + " = ?"
	queryUpdateUserStaffFlag = "UPDATE " + TableName + " SET " + StaffFlagColumnName + " = ? WHERE " + UserNameColumnName + " = ?"
	queryDeleteUser          = "DELETE FROM " + TableName + " WHERE " + UserNameColumnName + " = ?"
)

// CreateTableIfNotExists creates table `users`.
func CreateTableIfNotExists(databasePtr *sql.DB) error {
	_, createTableError := databasePtr.Exec("CREATE TABLE IF NOT EXISTS " + TableName +
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
		StaffFlagColumnName + "		BOOLEAN			NOT NULL," +
		"PRIMARY KEY(" + UserNameColumnName + ")," +
		"UNIQUE(" + MailColumnName + ")," +
		"UNIQUE(" + NicknameColumnName + ")" +
		");")
	return createTableError
}

// InsertUser inserts the given user to the database `users` table.
func InsertUser(user User, databasePtr *sql.DB) Status {
	return queryDatabase(databasePtr, queryInsertUser, user.Mail, user.Password, user.UserName, user.Nickname, user.Fname, user.Lname, user.Phone, user.Location, user.Money, user.Introduction, user.StaffFlag)
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
		scanError := databaseUsersTableRowsPtr.Scan(&user.Mail, &user.Password, &user.UserName, &user.Nickname, &user.Fname, &user.Lname, &user.Phone, &user.Location, &user.Money, &user.Introduction, &user.StaffFlag)
		if scanError != nil {
			return nil, util.StatusInternalServerError(getAllUsers, scanError)
		}
		users = append(users, user)
	}
	return users, util.StatusOK()
}

// UpdateUserPassword updates the given user's password to the database `users` table.
func UpdateUserPassword(userName string, userNewPassword string, databasePtr *sql.DB) Status {
	return queryDatabase(databasePtr, queryUpdateUserPassword, userNewPassword, userName)
}

// UpdateUserStaffFlag updates the given user's staff flag to the database `users` table.
func UpdateUserStaffFlag(userName string, userNewStaffFlag string, databasePtr *sql.DB) Status {
	return queryDatabase(databasePtr, queryUpdateUserStaffFlag, userNewStaffFlag, userName)
}

// DeleteUser deletes the given user from the database `users` table.
func DeleteUser(userName string, databasePtr *sql.DB) Status {
	return queryDatabase(databasePtr, queryDeleteUser, userName)
}

// IsExistingUser determines whether the given user is in the database `users` table.
func IsExistingUser(userName string, databasePtr *sql.DB) (bool, Status) {
	users, getStatus := getUsersByKeyColumn(queryGetUserByUserName, userName, databasePtr)
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
	return getUserByKeyColumn(queryGetUserByMail, mail, databasePtr)
}

// GetUserByUserName returns an user's information by the given user name.
func GetUserByUserName(userName string, databasePtr *sql.DB) (User, Status) {
	return getUserByKeyColumn(queryGetUserByUserName, userName, databasePtr)
}

func getUserByKeyColumn(queryGetUserByKeyColumn string, keyColumnValue string, databasePtr *sql.DB) (User, Status) {
	var dumpUser User
	users, getStatus := getUsersByKeyColumn(queryGetUserByKeyColumn, keyColumnValue, databasePtr)
	if !util.IsStatusOK(getStatus) {
		return dumpUser, getStatus
	}
	if len(users) != 1 {
		return dumpUser, util.StatusInternalServerError(getUserByKeyColumn, errors.New("Query 1 user but got "+strconv.Itoa(len(users))+" user(s) instead."))
	}
	return users[0], util.StatusOK()
}

func getUsersByKeyColumn(queryGetUserByKeyColumn string, keyColumnValue string, databasePtr *sql.DB) ([]User, Status) {
	queryRowsPtr, queryError := databasePtr.Query(queryGetUserByKeyColumn, keyColumnValue)
	if queryError != nil {
		return nil, util.StatusInternalServerError(getUsersByKeyColumn, queryError)
	}
	defer queryRowsPtr.Close()
	users, getStatus := getAllUsers(queryRowsPtr)
	if !util.IsStatusOK(getStatus) {
		return nil, getStatus
	}
	return users, util.StatusOK()
}
