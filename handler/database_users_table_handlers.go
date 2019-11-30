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

type password struct {
	Value string `json:"password" binding:"required"`
}

// RespondJsonOfAllUsersFromDatabaseUsersTableHandler responds to the client the json of all users from the database table 'users'.
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
		scanError := databaseUsersTableRowsPtr.Scan(&user.UserName, &user.Password)
		if scanError != nil {
			return nil, util.StatusInternalServerError(getAllUsers, scanError)
		}
		users = append(users, user)
	}
	return users, util.StatusOK()
}

// CreateUserToDatabaseUsersTableAndRespondJsonOfUserHandler creates the user given in the context to the database table 'users'.
// Also, it responds to the client the json of the given user.
func CreateUserToDatabaseUsersTableAndRespondJsonOfUserHandler(databasePtr *sql.DB) gin.HandlerFunc {
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
		getUser, getStatus := getUserFromDatabaseUsersTable(user.UserName, databasePtr)
		if !util.IsStatusOK(getStatus) {
			context.JSON(getStatus.HttpStatusCode, gin.H{util.JsonError: getStatus.ErrorMessage})
			return
		}
		context.JSON(http.StatusOK, getUser)
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
	prepareStatementPtr, prepareStatus := prepareInsertUserToDatabaseUsersTable(databasePtr)
	if !util.IsStatusOK(prepareStatus) {
		return prepareStatus
	}
	_, executeError := prepareStatementPtr.Exec(user.UserName, user.Password)
	if executeError != nil {
		return util.StatusInternalServerError(insertUserToDatabaseUsersTable, executeError)
	}
	return util.StatusOK()
}

func prepareInsertUserToDatabaseUsersTable(databasePtr *sql.DB) (*sql.Stmt, Status) {
	prepareStatementPtr, prepareError := databasePtr.Prepare("INSERT INTO " + DUTU.TableName + " VALUES(?, ?)")
	if prepareError != nil {
		return nil, util.StatusInternalServerError(prepareInsertUserToDatabaseUsersTable, prepareError)
	}
	return prepareStatementPtr, util.StatusOK()
}

// RespondJsonOfUserFromDatabaseUsersTableHandler responds to the client the json of the user given in the context parameter from the database table 'users'.
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

// UpdateUserPasswordInDatabaseUsersTableAndRespondJsonOfUserHandler updates the password of the user in the database table 'users' whose name is given in the context parameter and the requested JSON object.
// Also, it responds to the client the json of the given user.
func UpdateUserPasswordInDatabaseUsersTableAndRespondJsonOfUserHandler(databasePtr *sql.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		userName := context.Param(DUTU.UserNameColumnName)
		userNewPassword, getStatus := getPasswordFromContext(context)
		if !util.IsStatusOK(getStatus) {
			context.JSON(getStatus.HttpStatusCode, gin.H{util.JsonError: getStatus.ErrorMessage})
			return
		}
		var user User
		user.UserName = userName
		user.Password = userNewPassword.Value
		updateStatus := updateUserPasswordToDatabaseUsersTable(user, databasePtr)
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

func updateUserPasswordToDatabaseUsersTable(userOfNewPassword User, databasePtr *sql.DB) Status {
	prepareStatementPtr, prepareStatus := prepareUpdateUserPasswordToDatabaseUsersTable(databasePtr)
	if !util.IsStatusOK(prepareStatus) {
		return prepareStatus
	}
	_, executeError := prepareStatementPtr.Exec(userOfNewPassword.Password, userOfNewPassword.UserName)
	if executeError != nil {
		return util.StatusInternalServerError(updateUserPasswordToDatabaseUsersTable, executeError)
	}
	return util.StatusOK()
}

func prepareUpdateUserPasswordToDatabaseUsersTable(databasePtr *sql.DB) (*sql.Stmt, Status) {
	prepareStatementPtr, prepareError := databasePtr.Prepare("UPDATE " + DUTU.TableName + " SET " + DUTU.PasswordColumnName + " = ? WHERE " + DUTU.UserNameColumnName + " = ?")
	if prepareError != nil {
		return nil, util.StatusInternalServerError(prepareUpdateUserPasswordToDatabaseUsersTable, prepareError)
	}
	return prepareStatementPtr, util.StatusOK()
}

// DeleteUserFromDatabaseUsersTableAndRespondJsonOfUserNameHandler deletes the user whose name is given in the context parameter from the database table 'users'.
// Also, it responds to the client the json of the given user name.
func DeleteUserFromDatabaseUsersTableAndRespondJsonOfUserNameHandler(databasePtr *sql.DB) gin.HandlerFunc {
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
	prepareStatementPtr, prepareStatus := prepareDeleteUserFromDatabaseUsersTable(databasePtr)
	if !util.IsStatusOK(prepareStatus) {
		return prepareStatus
	}
	_, executeError := prepareStatementPtr.Exec(userName)
	if executeError != nil {
		return util.StatusInternalServerError(deleteUserFromDatabaseUsersTable, executeError)
	}
	return util.StatusOK()
}

func prepareDeleteUserFromDatabaseUsersTable(databasePtr *sql.DB) (*sql.Stmt, Status) {
	prepareStatementPtr, prepareError := databasePtr.Prepare("DELETE FROM " + DUTU.TableName + " WHERE " + DUTU.UserNameColumnName + " = ?")
	if prepareError != nil {
		return nil, util.StatusInternalServerError(prepareDeleteUserFromDatabaseUsersTable, prepareError)
	}
	return prepareStatementPtr, util.StatusOK()
}
