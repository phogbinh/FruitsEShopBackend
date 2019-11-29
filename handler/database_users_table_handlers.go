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

// ResponseJsonOfAllUsersFromDatabaseUsersTableHandler responses to the client the json of all users from the database table 'users'.
func ResponseJsonOfAllUsersFromDatabaseUsersTableHandler(databasePtr *sql.DB) gin.HandlerFunc {
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
		scanError := databaseUsersTableRowsPtr.Scan(&user.UserName, &user.Password)
		if scanError != nil {
			return nil, util.StatusInternalServerError(getAllUsers, scanError)
		}
		users = append(users, user)
	}
	return users, util.StatusOK()
}

// CreateUserToDatabaseUsersTableAndResponseJsonOfUserHandler creates the user given in the context to the database table 'users'.
// Also, it responses to the client the json of the given user.
func CreateUserToDatabaseUsersTableAndResponseJsonOfUserHandler(databasePtr *sql.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		user, getStatus := getUserFromContext(context)
		if !util.IsStatusOK(getStatus) {
			context.JSON(getStatus.HttpStatusCode, gin.H{util.JsonError: getStatus.ErrorMessage})
			return
		}
		insertStatus := insertUserToDatabaseUsersTable(user, databasePtr)
		if !util.IsStatusOK(insertStatus) {
			context.JSON(insertStatus.HttpStatusCode, gin.H{util.JsonError: insertStatus.ErrorMessage})
			return
		}
		context.JSON(http.StatusOK, user)
	}
}

func getUserFromContext(context *gin.Context) (User, Status) {
	var user User
	bindError := context.ShouldBindJSON(&user)
	if bindError != nil {
		return user, util.StatusBadRequest(getUserFromContext, bindError)
	}
	return user, util.StatusOK()
}

func insertUserToDatabaseUsersTable(user User, databasePtr *sql.DB) Status {
	prepareStatementPtr, prepareStatus := prepareInsertUserToDatabaseUsersTable(user, databasePtr)
	if !util.IsStatusOK(prepareStatus) {
		return prepareStatus
	}
	_, executeError := prepareStatementPtr.Exec(user.UserName, user.Password)
	if executeError != nil {
		return util.StatusInternalServerError(insertUserToDatabaseUsersTable, executeError)
	}
	return util.StatusOK()
}

func prepareInsertUserToDatabaseUsersTable(user User, databasePtr *sql.DB) (*sql.Stmt, Status) {
	prepareStatementPtr, prepareError := databasePtr.Prepare("INSERT INTO " + DUTU.TableName + " VALUES(?, ?)")
	if prepareError != nil {
		return nil, util.StatusInternalServerError(prepareInsertUserToDatabaseUsersTable, prepareError)
	}
	return prepareStatementPtr, util.StatusOK()
}

// ResponseJsonOfUserFromDatabaseUsersTableHandler responses to the client the json of the user given in the context parameter from the database table 'users'.
func ResponseJsonOfUserFromDatabaseUsersTableHandler(databasePtr *sql.DB) gin.HandlerFunc {
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

// UpdateUserPasswordInDatabaseUsersTableAndResponseJsonOfUserHandler updates the password of the user in the database table 'users' whose name is given in the context parameter and the requested JSON object.
// Also, it responses to the client the json of the given user.
func UpdateUserPasswordInDatabaseUsersTableAndResponseJsonOfUserHandler(databasePtr *sql.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		userName := context.Param(DUTU.UserNameColumnName)
		newPasswordUser, getStatus := getUserFromContext(context)
		if !util.IsStatusOK(getStatus) {
			context.JSON(getStatus.HttpStatusCode, gin.H{util.JsonError: getStatus.ErrorMessage})
			return
		}
		if userName != newPasswordUser.UserName {
			context.JSON(http.StatusBadRequest, gin.H{util.JsonError: "The user name given in the context parameter - " + userName + " - does not match the user name provided by the requested JSON object - " + newPasswordUser.UserName + "."})
			return
		}
		updateStatus := updateUserPasswordToDatabaseUsersTable(newPasswordUser, databasePtr)
		if !util.IsStatusOK(updateStatus) {
			context.JSON(updateStatus.HttpStatusCode, gin.H{util.JsonError: updateStatus.ErrorMessage})
			return
		}
		context.JSON(http.StatusOK, newPasswordUser)
	}
}

func updateUserPasswordToDatabaseUsersTable(userOfNewPassword User, databasePtr *sql.DB) Status {
	prepareStatementPtr, prepareStatus := prepareUpdateUserPasswordToDatabaseUsersTable(userOfNewPassword, databasePtr)
	if !util.IsStatusOK(prepareStatus) {
		return prepareStatus
	}
	_, executeError := prepareStatementPtr.Exec(userOfNewPassword.Password, userOfNewPassword.UserName)
	if executeError != nil {
		return util.StatusInternalServerError(updateUserPasswordToDatabaseUsersTable, executeError)
	}
	return util.StatusOK()
}

func prepareUpdateUserPasswordToDatabaseUsersTable(userOfNewPassword User, databasePtr *sql.DB) (*sql.Stmt, Status) {
	prepareStatementPtr, prepareError := databasePtr.Prepare("UPDATE " + DUTU.TableName + " SET " + DUTU.PasswordColumnName + " = ? WHERE " + DUTU.UserNameColumnName + " = ?")
	if prepareError != nil {
		return nil, util.StatusInternalServerError(prepareUpdateUserPasswordToDatabaseUsersTable, prepareError)
	}
	return prepareStatementPtr, util.StatusOK()
}

// DeleteUserFromDatabaseUsersTableAndResponseJsonOfUserNameHandler deletes the user whose name is given in the context parameter from the database table 'users'.
// Also, it responses to the client the json of the given user name.
func DeleteUserFromDatabaseUsersTableAndResponseJsonOfUserNameHandler(databasePtr *sql.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		userName := context.Param(DUTU.UserNameColumnName)
		deleteStatus := deleteUserFromDatabaseUsersTable(userName, databasePtr)
		if !util.IsStatusOK(deleteStatus) {
			context.JSON(deleteStatus.HttpStatusCode, gin.H{util.JsonError: deleteStatus.ErrorMessage})
			return
		}
		context.JSON(http.StatusOK, gin.H{DUTU.UserNameColumnName: userName})
	}
}

func deleteUserFromDatabaseUsersTable(userName string, databasePtr *sql.DB) Status {
	prepareStatementPtr, prepareStatus := prepareDeleteUserFromDatabaseUsersTable(userName, databasePtr)
	if !util.IsStatusOK(prepareStatus) {
		return prepareStatus
	}
	_, executeError := prepareStatementPtr.Exec(userName)
	if executeError != nil {
		return util.StatusInternalServerError(deleteUserFromDatabaseUsersTable, executeError)
	}
	return util.StatusOK()
}

func prepareDeleteUserFromDatabaseUsersTable(userName string, databasePtr *sql.DB) (*sql.Stmt, Status) {
	prepareStatementPtr, prepareError := databasePtr.Prepare("DELETE FROM " + DUTU.TableName + " WHERE " + DUTU.UserNameColumnName + " = ?")
	if prepareError != nil {
		return nil, util.StatusInternalServerError(prepareDeleteUserFromDatabaseUsersTable, prepareError)
	}
	return prepareStatementPtr, util.StatusOK()
}
