package handler

import (
	"database/sql"
	"net/http"
	"strconv"

	DUTU "backend/database_users_table_util"
	. "backend/model"
	"backend/util"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

const (
	noError = ""
)

// ResponseJsonOfAllUsersFromDatabaseUsersTableHandler responses to the client the json of all users from the database table 'users'.
func ResponseJsonOfAllUsersFromDatabaseUsersTableHandler(databasePtr *sql.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		users, status := getAllUsersFromDatabaseUsersTable(databasePtr)
		if status.HttpStatusCode != http.StatusOK {
			context.String(status.HttpStatusCode, status.ErrorMessage)
			return
		}
		context.JSON(http.StatusOK, users)
	}
}

func getAllUsersFromDatabaseUsersTable(databasePtr *sql.DB) ([]User, Status) {
	queryRowsPtr, queryError := databasePtr.Query("SELECT * FROM " + DUTU.TableName)
	if queryError != nil {
		return nil, Status{
			HttpStatusCode: http.StatusInternalServerError,
			ErrorMessage:   util.GetErrorMessageHeaderContainingFunctionName(getAllUsersFromDatabaseUsersTable) + queryError.Error()}
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
			return nil, Status{
				HttpStatusCode: http.StatusInternalServerError,
				ErrorMessage:   util.GetErrorMessageHeaderContainingFunctionName(getAllUsers) + scanError.Error()}
		}
		users = append(users, user)
	}
	return users, Status{
		HttpStatusCode: http.StatusOK,
		ErrorMessage:   noError}
}

// CreateUserToDatabaseUsersTableAndResponseJsonOfUserHandler creates the user given in the context to the database table 'users'.
// Also, it responses to the client the json of the given user.
func CreateUserToDatabaseUsersTableAndResponseJsonOfUserHandler(databasePtr *sql.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		user, getStatus := getUserFromContext(context)
		if getStatus.HttpStatusCode != http.StatusOK {
			context.String(getStatus.HttpStatusCode, getStatus.ErrorMessage)
			return
		}
		insertStatus := insertUserToDatabaseUsersTable(user, databasePtr)
		if insertStatus.HttpStatusCode != http.StatusOK {
			context.String(insertStatus.HttpStatusCode, insertStatus.ErrorMessage)
			return
		}
		context.JSON(http.StatusOK, user)
	}
}

func getUserFromContext(context *gin.Context) (User, Status) {
	var user User
	bindError := context.ShouldBindJSON(&user)
	if bindError != nil {
		return user, Status{
			HttpStatusCode: http.StatusBadRequest,
			ErrorMessage:   util.GetErrorMessageHeaderContainingFunctionName(getUserFromContext) + bindError.Error()}
	}
	return user, Status{
		HttpStatusCode: http.StatusOK,
		ErrorMessage:   noError}
}

func insertUserToDatabaseUsersTable(user User, databasePtr *sql.DB) Status {
	prepareStatementPtr, prepareStatus := prepareInsertUserToDatabaseUsersTable(user, databasePtr)
	if prepareStatus.HttpStatusCode != http.StatusOK {
		return prepareStatus
	}
	_, executeError := prepareStatementPtr.Exec(user.UserName, user.Password)
	if executeError != nil {
		return Status{
			HttpStatusCode: http.StatusInternalServerError,
			ErrorMessage:   util.GetErrorMessageHeaderContainingFunctionName(insertUserToDatabaseUsersTable) + executeError.Error()}
	}
	return Status{
		HttpStatusCode: http.StatusOK,
		ErrorMessage:   noError}
}

func prepareInsertUserToDatabaseUsersTable(user User, databasePtr *sql.DB) (*sql.Stmt, Status) {
	prepareStatementPtr, prepareError := databasePtr.Prepare("INSERT INTO " + DUTU.TableName + " VALUES(?, ?)")
	if prepareError != nil {
		return nil, Status{
			HttpStatusCode: http.StatusInternalServerError,
			ErrorMessage:   util.GetErrorMessageHeaderContainingFunctionName(prepareInsertUserToDatabaseUsersTable) + prepareError.Error()}
	}
	return prepareStatementPtr, Status{
		HttpStatusCode: http.StatusOK,
		ErrorMessage:   noError}
}

// ResponseJsonOfUserFromDatabaseUsersTableHandler responses to the client the json of the user given in the context parameter from the database table 'users'.
func ResponseJsonOfUserFromDatabaseUsersTableHandler(databasePtr *sql.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		userName := context.Param(DUTU.UserNameColumnName)
		user, status := getUserFromDatabaseUsersTable(userName, databasePtr)
		if status.HttpStatusCode != http.StatusOK {
			context.String(status.HttpStatusCode, status.ErrorMessage)
			return
		}
		context.JSON(http.StatusOK, user)
	}
}

func getUserFromDatabaseUsersTable(userName string, databasePtr *sql.DB) (User, Status) {
	var dumpUser User
	queryRowsPtr, queryStatus := getUserQueryRowsPtrFromDatabaseUsersTable(userName, databasePtr)
	if queryStatus.HttpStatusCode != http.StatusOK {
		return dumpUser, queryStatus
	}
	defer queryRowsPtr.Close()
	users, getStatus := getAllUsers(queryRowsPtr)
	if getStatus.HttpStatusCode != http.StatusOK {
		return dumpUser, getStatus
	}
	if len(users) != 1 {
		return dumpUser, Status{
			HttpStatusCode: http.StatusInternalServerError,
			ErrorMessage:   util.GetErrorMessageHeaderContainingFunctionName(getUserFromDatabaseUsersTable) + strconv.Itoa(len(users)) + " user(s)."}
	}
	return users[0], Status{
		HttpStatusCode: http.StatusOK,
		ErrorMessage:   noError}
}

func getUserQueryRowsPtrFromDatabaseUsersTable(userName string, databasePtr *sql.DB) (*sql.Rows, Status) {
	queryRowsPtr, queryError := databasePtr.Query("SELECT * FROM "+DUTU.TableName+" WHERE "+DUTU.UserNameColumnName+" = ?", userName)
	if queryError != nil {
		return nil, Status{
			HttpStatusCode: http.StatusInternalServerError,
			ErrorMessage:   util.GetErrorMessageHeaderContainingFunctionName(getUserQueryRowsPtrFromDatabaseUsersTable) + queryError.Error()}
	}
	return queryRowsPtr, Status{
		HttpStatusCode: http.StatusOK,
		ErrorMessage:   noError}
}

// UpdateUserPasswordInDatabaseUsersTableAndResponseJsonOfUserHandler updates the password of the user in the database table 'users' whose name is given in the context parameter and the requested JSON object.
// Also, it responses to the client the json of the given user.
func UpdateUserPasswordInDatabaseUsersTableAndResponseJsonOfUserHandler(databasePtr *sql.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		userName := context.Param(DUTU.UserNameColumnName)
		newPasswordUser, getStatus := getUserFromContext(context)
		if getStatus.HttpStatusCode != http.StatusOK {
			context.String(getStatus.HttpStatusCode, getStatus.ErrorMessage)
			return
		}
		if userName != newPasswordUser.UserName {
			context.String(http.StatusBadRequest, "The user name given in the context parameter - "+userName+" - does not match the user name provided by the requested JSON object - "+newPasswordUser.UserName+".")
			return
		}
		updateStatus := updateUserPasswordToDatabaseUsersTable(newPasswordUser, databasePtr)
		if updateStatus.HttpStatusCode != http.StatusOK {
			context.String(updateStatus.HttpStatusCode, updateStatus.ErrorMessage)
			return
		}
		context.JSON(http.StatusOK, newPasswordUser)
	}
}

func updateUserPasswordToDatabaseUsersTable(userOfNewPassword User, databasePtr *sql.DB) Status {
	prepareStatementPtr, prepareStatus := prepareUpdateUserPasswordToDatabaseUsersTable(userOfNewPassword, databasePtr)
	if prepareStatus.HttpStatusCode != http.StatusOK {
		return prepareStatus
	}
	_, executeError := prepareStatementPtr.Exec(userOfNewPassword.Password, userOfNewPassword.UserName)
	if executeError != nil {
		return Status{
			HttpStatusCode: http.StatusInternalServerError,
			ErrorMessage:   util.GetErrorMessageHeaderContainingFunctionName(updateUserPasswordToDatabaseUsersTable) + executeError.Error()}
	}
	return Status{
		HttpStatusCode: http.StatusOK,
		ErrorMessage:   noError}
}

func prepareUpdateUserPasswordToDatabaseUsersTable(userOfNewPassword User, databasePtr *sql.DB) (*sql.Stmt, Status) {
	prepareStatementPtr, prepareError := databasePtr.Prepare("UPDATE " + DUTU.TableName + " SET " + DUTU.PasswordColumnName + " = ? WHERE " + DUTU.UserNameColumnName + " = ?")
	if prepareError != nil {
		return nil, Status{
			HttpStatusCode: http.StatusInternalServerError,
			ErrorMessage:   util.GetErrorMessageHeaderContainingFunctionName(prepareUpdateUserPasswordToDatabaseUsersTable) + prepareError.Error()}
	}
	return prepareStatementPtr, Status{
		HttpStatusCode: http.StatusOK,
		ErrorMessage:   noError}
}

// DeleteUserFromDatabaseUsersTableAndResponseJsonOfUserNameHandler deletes the user whose name is given in the context parameter from the database table 'users'.
// Also, it responses to the client the json of the given user name.
func DeleteUserFromDatabaseUsersTableAndResponseJsonOfUserNameHandler(databasePtr *sql.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		userName := context.Param(DUTU.UserNameColumnName)
		deleteStatus := deleteUserFromDatabaseUsersTable(userName, databasePtr)
		if deleteStatus.HttpStatusCode != http.StatusOK {
			context.String(deleteStatus.HttpStatusCode, deleteStatus.ErrorMessage)
			return
		}
		context.JSON(http.StatusOK, gin.H{DUTU.UserNameColumnName: userName})
	}
}

func deleteUserFromDatabaseUsersTable(userName string, databasePtr *sql.DB) Status {
	prepareStatementPtr, prepareStatus := prepareDeleteUserFromDatabaseUsersTable(userName, databasePtr)
	if prepareStatus.HttpStatusCode != http.StatusOK {
		return prepareStatus
	}
	_, executeError := prepareStatementPtr.Exec(userName)
	if executeError != nil {
		return Status{
			HttpStatusCode: http.StatusInternalServerError,
			ErrorMessage:   util.GetErrorMessageHeaderContainingFunctionName(deleteUserFromDatabaseUsersTable) + executeError.Error()}
	}
	return Status{
		HttpStatusCode: http.StatusOK,
		ErrorMessage:   noError}
}

func prepareDeleteUserFromDatabaseUsersTable(userName string, databasePtr *sql.DB) (*sql.Stmt, Status) {
	prepareStatementPtr, prepareError := databasePtr.Prepare("DELETE FROM " + DUTU.TableName + " WHERE " + DUTU.UserNameColumnName + " = ?")
	if prepareError != nil {
		return nil, Status{
			HttpStatusCode: http.StatusInternalServerError,
			ErrorMessage:   util.GetErrorMessageHeaderContainingFunctionName(prepareDeleteUserFromDatabaseUsersTable) + prepareError.Error()}
	}
	return prepareStatementPtr, Status{
		HttpStatusCode: http.StatusOK,
		ErrorMessage:   noError}
}
