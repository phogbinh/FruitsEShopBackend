package handler

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	DUTU "backend/database_users_table_util"
	. "backend/model"
	"backend/util"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

const (
	queryInsertUser         = "INSERT INTO " + DUTU.TableName + " VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	queryGetAllUsers        = "SELECT * FROM " + DUTU.TableName
	queryGetUserByUserName  = "SELECT * FROM " + DUTU.TableName + " WHERE " + DUTU.UserNameColumnName + " = ?"
	queryGetUserByMail      = "SELECT * FROM " + DUTU.TableName + " WHERE " + DUTU.MailColumnName + " = ?"
	queryUpdateUserPassword = "UPDATE " + DUTU.TableName + " SET " + DUTU.PasswordColumnName + " = ? WHERE " + DUTU.UserNameColumnName + " = ?"
	queryDeleteUser         = "DELETE FROM " + DUTU.TableName + " WHERE " + DUTU.UserNameColumnName + " = ?"
)

type password struct {
	Value string `json:"password" binding:"required"`
}

// CreateUserToDatabaseUsersTableAndRespondJsonOfUserHandler creates an user and responds the user's information.
func CreateUserToDatabaseUsersTableAndRespondJsonOfUserHandler(databasePtr *sql.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		user, getStatus := getUserFromRequest(context)
		if !util.IsStatusOK(getStatus) {
			context.JSON(getStatus.HttpStatusCode, gin.H{util.JsonError: getStatus.ErrorMessage})
			return
		}
		insertStatus := InsertUser(user, databasePtr)
		if !util.IsStatusOK(insertStatus) {
			context.JSON(insertStatus.HttpStatusCode, gin.H{util.JsonError: insertStatus.ErrorMessage})
			return
		}
		getUser, getStatus := getUserByKeyColumnFromDatabaseUsersTable(queryGetUserByUserName, user.UserName, databasePtr)
		if !util.IsStatusOK(getStatus) {
			context.JSON(getStatus.HttpStatusCode, gin.H{util.JsonError: getStatus.ErrorMessage})
			return
		}
		context.JSON(http.StatusOK, getUser)
	}
}

func getUserFromRequest(context *gin.Context) (User, Status) {
	var user User
	bindError := context.ShouldBindJSON(&user)
	if bindError != nil {
		return user, util.StatusBadRequest(getUserFromRequest, bindError)
	}
	return user, util.StatusOK()
}

// InsertUser inserts the given user to the database `users` table.
func InsertUser(user User, databasePtr *sql.DB) Status {
	return queryDatabase(databasePtr, queryInsertUser, user.Mail, user.Password, user.UserName, user.Nickname, user.Fname, user.Lname, user.Phone, user.Location, user.Money, user.Introduction)
}

// RespondJsonOfAllUsersFromDatabaseUsersTableHandler responds all users' information.
func RespondJsonOfAllUsersFromDatabaseUsersTableHandler(databasePtr *sql.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		users, status := GetAllUsers(databasePtr)
		if !util.IsStatusOK(status) {
			context.JSON(status.HttpStatusCode, gin.H{util.JsonError: status.ErrorMessage})
			return
		}
		context.JSON(http.StatusOK, users)
	}
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

// RespondJsonOfUserByUserNameFromDatabaseUsersTableHandler responds an user's information.
func RespondJsonOfUserByUserNameFromDatabaseUsersTableHandler(databasePtr *sql.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		userName := context.Param(DUTU.UserNameColumnName)
		user, status := getUserByKeyColumnFromDatabaseUsersTable(queryGetUserByUserName, userName, databasePtr)
		if !util.IsStatusOK(status) {
			context.JSON(status.HttpStatusCode, gin.H{util.JsonError: status.ErrorMessage})
			return
		}
		context.JSON(http.StatusOK, user)
	}
}

// RespondJsonOfUserByMailFromDatabaseUsersTableHandler responds an user's information by the given mail.
func RespondJsonOfUserByMailFromDatabaseUsersTableHandler(databasePtr *sql.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		mail := context.Query(DUTU.MailColumnName)
		user, status := getUserByKeyColumnFromDatabaseUsersTable(queryGetUserByMail, mail, databasePtr)
		if !util.IsStatusOK(status) {
			context.JSON(status.HttpStatusCode, gin.H{util.JsonError: status.ErrorMessage})
			return
		}
		context.JSON(http.StatusOK, user)
	}
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

// UpdateUserPasswordInDatabaseUsersTableAndRespondJsonOfUserHandler updates an user's password and responds the user's information.
func UpdateUserPasswordInDatabaseUsersTableAndRespondJsonOfUserHandler(databasePtr *sql.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		userName := context.Param(DUTU.UserNameColumnName)
		userNewPassword, getStatus := getPasswordFromRequest(context)
		if !util.IsStatusOK(getStatus) {
			context.JSON(getStatus.HttpStatusCode, gin.H{util.JsonError: getStatus.ErrorMessage})
			return
		}
		updateStatus := UpdateUserPassword(userName, userNewPassword.Value, databasePtr)
		if !util.IsStatusOK(updateStatus) {
			context.JSON(updateStatus.HttpStatusCode, gin.H{util.JsonError: updateStatus.ErrorMessage})
			return
		}
		getUser, getStatus := getUserByKeyColumnFromDatabaseUsersTable(queryGetUserByUserName, userName, databasePtr)
		if !util.IsStatusOK(getStatus) {
			context.JSON(getStatus.HttpStatusCode, gin.H{util.JsonError: getStatus.ErrorMessage})
			return
		}
		context.JSON(http.StatusOK, getUser)
	}
}

func getPasswordFromRequest(context *gin.Context) (password, Status) {
	var pwd password
	bindError := context.ShouldBindJSON(&pwd)
	if bindError != nil {
		return pwd, util.StatusBadRequest(getPasswordFromRequest, bindError)
	}
	return pwd, util.StatusOK()
}

// UpdateUserPassword updates the given user's password to the database `users` table.
func UpdateUserPassword(userName string, userNewPassword string, databasePtr *sql.DB) Status {
	return queryDatabase(databasePtr, queryUpdateUserPassword, userNewPassword, userName)
}

// DeleteUserFromDatabaseUsersTable deletes an user.
func DeleteUserFromDatabaseUsersTable(databasePtr *sql.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		userName := context.Param(DUTU.UserNameColumnName)
		deleteStatus := DeleteUser(userName, databasePtr)
		if !util.IsStatusOK(deleteStatus) {
			context.JSON(deleteStatus.HttpStatusCode, gin.H{util.JsonError: deleteStatus.ErrorMessage})
			return
		}
		isExistingUser, getStatus := IsExistingUser(userName, databasePtr)
		if !util.IsStatusOK(getStatus) {
			context.JSON(getStatus.HttpStatusCode, gin.H{util.JsonError: getStatus.ErrorMessage})
			return
		}
		if isExistingUser {
			context.JSON(http.StatusInternalServerError, gin.H{util.JsonError: "User still exists."})
			return
		}
		context.Status(http.StatusOK)
	}
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
