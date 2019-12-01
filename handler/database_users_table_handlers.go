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
	queryUpdateUserPassword = "UPDATE " + DUTU.TableName + " SET " + DUTU.PasswordColumnName + " = ? WHERE " + DUTU.UserNameColumnName + " = ?"
	queryDeleteUser         = "DELETE FROM " + DUTU.TableName + " WHERE " + DUTU.UserNameColumnName + " = ?"
)

type password struct {
	Value string `json:"password" binding:"required"`
}

func CreateUserToDatabaseUsersTableAndRespondJsonOfUserHandler(databasePtr *sql.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		user, getStatus := getUserFromRequest(context)
		if !util.IsStatusOK(getStatus) {
			context.JSON(getStatus.HttpStatusCode, gin.H{util.JsonError: getStatus.ErrorMessage})
			return
		}
		insertStatus := insertUserToDatabaseUsersTable(user, databasePtr)
		if !util.IsStatusOK(insertStatus) {
			context.JSON(insertStatus.HttpStatusCode, gin.H{util.JsonError: insertStatus.ErrorMessage})
			return
		}
		getUser, getStatus := getUserFromDatabaseUsersTable(user.UserName, databasePtr)
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

func insertUserToDatabaseUsersTable(user User, databasePtr *sql.DB) Status {
	return queryDatabase(databasePtr, queryInsertUser, user.Mail, user.Password, user.UserName, user.Nickname, user.Fname, user.Lname, user.Phone, user.Location, user.Money, user.Introduction)
}

func RespondJsonOfAllUsersFromDatabaseUsersTableHandler(databasePtr *sql.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		users, status := getAllUsersFromDatabaseUsersTable(databasePtr)
		if !util.IsStatusOK(status) {
			context.JSON(status.HttpStatusCode, gin.H{util.JsonError: status.ErrorMessage})
			return
		}
		context.JSON(http.StatusOK, users)
	}
}

func getAllUsersFromDatabaseUsersTable(databasePtr *sql.DB) ([]User, Status) {
	queryRowsPtr, queryError := databasePtr.Query("SELECT * FROM " + DUTU.TableName)
	if queryError != nil {
		return nil, util.StatusInternalServerError(getAllUsersFromDatabaseUsersTable, queryError)
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

func RespondJsonOfUserFromDatabaseUsersTableHandler(databasePtr *sql.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		userName := context.Param(DUTU.UserNameColumnName)
		user, status := getUserFromDatabaseUsersTable(userName, databasePtr)
		if !util.IsStatusOK(status) {
			context.JSON(status.HttpStatusCode, gin.H{util.JsonError: status.ErrorMessage})
			return
		}
		context.JSON(http.StatusOK, user)
	}
}

func getUserFromDatabaseUsersTable(userName string, databasePtr *sql.DB) (User, Status) {
	var dumpUser User
	queryRowsPtr, queryStatus := getUserQueryRowsPtrFromDatabaseUsersTable(userName, databasePtr)
	if !util.IsStatusOK(queryStatus) {
		return dumpUser, queryStatus
	}
	defer queryRowsPtr.Close()
	users, getStatus := getAllUsers(queryRowsPtr)
	if !util.IsStatusOK(getStatus) {
		return dumpUser, getStatus
	}
	if len(users) != 1 {
		return dumpUser, util.StatusInternalServerError(getUserFromDatabaseUsersTable, errors.New("Query 1 user but got "+strconv.Itoa(len(users))+" user(s) instead."))
	}
	return users[0], util.StatusOK()
}

func getUserQueryRowsPtrFromDatabaseUsersTable(userName string, databasePtr *sql.DB) (*sql.Rows, Status) {
	queryRowsPtr, queryError := databasePtr.Query("SELECT * FROM "+DUTU.TableName+" WHERE "+DUTU.UserNameColumnName+" = ?", userName)
	if queryError != nil {
		return nil, util.StatusInternalServerError(getUserQueryRowsPtrFromDatabaseUsersTable, queryError)
	}
	return queryRowsPtr, util.StatusOK()
}

func UpdateUserPasswordInDatabaseUsersTableAndRespondJsonOfUserHandler(databasePtr *sql.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		userName := context.Param(DUTU.UserNameColumnName)
		userNewPassword, getStatus := getPasswordFromContext(context)
		if !util.IsStatusOK(getStatus) {
			context.JSON(getStatus.HttpStatusCode, gin.H{util.JsonError: getStatus.ErrorMessage})
			return
		}
		updateStatus := updateUserPasswordToDatabaseUsersTable(userName, userNewPassword.Value, databasePtr)
		if !util.IsStatusOK(updateStatus) {
			context.JSON(updateStatus.HttpStatusCode, gin.H{util.JsonError: updateStatus.ErrorMessage})
			return
		}
		getUser, getStatus := getUserFromDatabaseUsersTable(userName, databasePtr)
		if !util.IsStatusOK(getStatus) {
			context.JSON(getStatus.HttpStatusCode, gin.H{util.JsonError: getStatus.ErrorMessage})
			return
		}
		context.JSON(http.StatusOK, getUser)
	}
}

func getPasswordFromContext(context *gin.Context) (password, Status) {
	var pwd password
	bindError := context.ShouldBindJSON(&pwd)
	if bindError != nil {
		return pwd, util.StatusBadRequest(getPasswordFromContext, bindError)
	}
	return pwd, util.StatusOK()
}

func updateUserPasswordToDatabaseUsersTable(userName string, userNewPassword string, databasePtr *sql.DB) Status {
	return queryDatabase(databasePtr, queryUpdateUserPassword, userNewPassword, userName)
}

func DeleteUserFromDatabaseUsersTableAndRespondJsonOfUserHandler(databasePtr *sql.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		userName := context.Param(DUTU.UserNameColumnName)
		deleteStatus := deleteUserFromDatabaseUsersTable(userName, databasePtr)
		if !util.IsStatusOK(deleteStatus) {
			context.JSON(deleteStatus.HttpStatusCode, gin.H{util.JsonError: deleteStatus.ErrorMessage})
			return
		}
		getUser, getStatus := getUserFromDatabaseUsersTable(userName, databasePtr)
		if !util.IsStatusOK(getStatus) {
			context.JSON(getStatus.HttpStatusCode, gin.H{util.JsonError: getStatus.ErrorMessage})
			return
		}
		context.JSON(http.StatusOK, getUser)
	}
}

func deleteUserFromDatabaseUsersTable(userName string, databasePtr *sql.DB) Status {
	return queryDatabase(databasePtr, queryDeleteUser, userName)
}

func queryDatabase(databasePtr *sql.DB, query string, executeArguments ...interface{}) Status {
	prepareStatementPtr, prepareError := databasePtr.Prepare(query)
	if prepareError != nil {
		return util.StatusInternalServerError(queryDatabase, prepareError)
	}
	_, executeError := prepareStatementPtr.Exec(executeArguments)
	if executeError != nil {
		return util.StatusInternalServerError(queryDatabase, executeError)
	}
	return util.StatusOK()
}
