package database_users_table_util

import (
	"database/sql"

	"backend/database_util"
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
	queryCreateTable = "CREATE TABLE IF NOT EXISTS " + TableName +
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
		");"
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
	return database_util.CreateTableIfNotExists(databasePtr, queryCreateTable)
}

// InsertUser inserts the given user to the database `users` table.
func InsertUser(user User, databasePtr *sql.DB) Status {
	return database_util.PrepareThenExecuteQuery(databasePtr, queryInsertUser, user.Mail, user.Password, user.UserName, user.Nickname, user.Fname, user.Lname, user.Phone, user.Location, user.Money, user.Introduction, user.StaffFlag)
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

// UpdateUserPassword updates the given user's password to the database `users` table.
func UpdateUserPassword(userName string, userNewPassword string, databasePtr *sql.DB) Status {
	return database_util.PrepareThenExecuteQuery(databasePtr, queryUpdateUserPassword, userNewPassword, userName)
}

// UpdateUserStaffFlag updates the given user's staff flag to the database `users` table.
func UpdateUserStaffFlag(userName string, userNewStaffFlag string, databasePtr *sql.DB) Status {
	return database_util.PrepareThenExecuteQuery(databasePtr, queryUpdateUserStaffFlag, userNewStaffFlag, userName)
}

// DeleteUser deletes the given user from the database `users` table.
func DeleteUser(userName string, databasePtr *sql.DB) Status {
	return database_util.PrepareThenExecuteQuery(databasePtr, queryDeleteUser, userName)
}

// IsExistingUser determines whether the given user is in the database `users` table.
func IsExistingUser(userName string, databasePtr *sql.DB) (bool, Status) {
	users, getStatus := getUsersByKeyColumn(queryGetUserByUserName, userName, databasePtr)
	if !util.IsStatusOK(getStatus) {
		return false, getStatus
	}
	return len(users) > 0, util.StatusOK()
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
	object, getStatus := database_util.GetObjectByKeyColumn(queryGetUserByKeyColumn, keyColumnValue, databasePtr)
	if !util.IsStatusOK(getStatus) {
		return *new(User), getStatus
	}
	return getUser(object)
}

func getUsersByKeyColumn(queryGetUserByKeyColumn string, keyColumnValue string, databasePtr *sql.DB) ([]User, Status) {
	objects, getStatus := database_util.GetObjectsByKeyColumn(queryGetUserByKeyColumn, keyColumnValue, databasePtr)
	if !util.IsStatusOK(getStatus) {
		return nil, getStatus
	}
	return getUsers(objects)
}

func getAllUsers(databaseUsersTableRowsPtr *sql.Rows) ([]User, Status) {
	objects, getStatus := database_util.GetAllObjects(databaseUsersTableRowsPtr)
	if !util.IsStatusOK(getStatus) {
		return nil, getStatus
	}
	return getUsers(objects)
}

func getUsers(objects [][]interface{}) ([]User, Status) {
	var users []User
	for _, object := range objects {
		user, getStatus := getUser(object)
		if !util.IsStatusOK(getStatus) {
			return nil, getStatus
		}
		users = append(users, user)
	}
	return users, util.StatusOK()
}

func getUser(object []interface{}) (User, Status) {
	rawBytesList, getStatus := database_util.GetRawBytesList(object)
	if !util.IsStatusOK(getStatus) {
		return *new(User), getStatus
	}
	var user User
	user.Mail = string(rawBytesList[0])
	user.Password = string(rawBytesList[1])
	user.UserName = string(rawBytesList[2])
	user.Nickname = string(rawBytesList[3])
	user.Fname = string(rawBytesList[4])
	user.Lname = string(rawBytesList[5])
	user.Phone = string(rawBytesList[6])
	user.Location = string(rawBytesList[7])
	user.Money = string(rawBytesList[8])
	user.Introduction = string(rawBytesList[9])
	user.StaffFlag = string(rawBytesList[10])
	return user, util.StatusOK()
}
