package handler

import (
	"database/sql"
	"net/http"
	"strconv"

	DUTU "backend/database_users_table_util"
	. "backend/model"
	"backend/symbolutil"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

const (
	noError                                               = ""
	errorText                                             = "Error "
	errorDatabaseTableText                                = " the database table '" + DUTU.TableName + "'"
	errorSelectGetAllUsersFromDatabaseUsersTable          = errorText + "selecting all users from" + errorDatabaseTableText + symbolutil.ColonSpace
	errorScanGetAllUsersFromDatabaseUsersTableRowsPointer = errorText + "scanning all users from" + errorDatabaseTableText + "'s rows pointer" + symbolutil.ColonSpace
	errorGetUserFromContext                               = errorText + "getting user from context" + symbolutil.ColonSpace
	errorPrepareInsertUserToDatabaseUsersTable            = errorText + "preparing to insert user to" + errorDatabaseTableText + symbolutil.ColonSpace
	errorInsertUserToDatabaseUsersTable                   = errorText + "inserting user to" + errorDatabaseTableText + symbolutil.ColonSpace
	errorSelectGetUserFromDatabaseUsersTable              = errorText + "selecting an user from" + errorDatabaseTableText + symbolutil.ColonSpace
	errorGetManyUsersGetUserFromDatabaseUsersTable        = errorText + "want to get one but got none or many users from" + errorDatabaseTableText + symbolutil.ColonSpace
	errorPrepareUpdateUserPasswordToDatabaseUsersTable    = errorText + "preparing to update user password to" + errorDatabaseTableText + symbolutil.ColonSpace
	errorUpdateUserPasswordToDatabaseUsersTable           = errorText + "updating user password to" + errorDatabaseTableText + symbolutil.ColonSpace
	errorPrepareDeleteUserFromDatabaseUsersTable          = errorText + "preparing to delete user to" + errorDatabaseTableText + symbolutil.ColonSpace
	errorDeleteUserFromDatabaseUsersTable                 = errorText + "deleting user to" + errorDatabaseTableText + symbolutil.ColonSpace
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
	selectRowsPtr, selectError := databasePtr.Query("SELECT * FROM " + DUTU.TableName)
	if selectError != nil {
		return nil, Status{
			HttpStatusCode: http.StatusInternalServerError,
			ErrorMessage:   errorSelectGetAllUsersFromDatabaseUsersTable + selectError.Error()}
	}
	defer selectRowsPtr.Close()
	return getAllUsers(selectRowsPtr)
}

func getAllUsers(databaseUsersTableRowsPtr *sql.Rows) ([]User, Status) {
	var users []User
	for databaseUsersTableRowsPtr.Next() {
		var user User
		scanError := databaseUsersTableRowsPtr.Scan(&user.UserName, &user.Password)
		if scanError != nil {
			return nil, Status{
				HttpStatusCode: http.StatusInternalServerError,
				ErrorMessage:   errorScanGetAllUsersFromDatabaseUsersTableRowsPointer + scanError.Error()}
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
			ErrorMessage:   errorGetUserFromContext + bindError.Error()}
	}
	return user, Status{
		HttpStatusCode: http.StatusOK,
		ErrorMessage:   noError}
}

func insertUserToDatabaseUsersTable(user User, databasePtr *sql.DB) Status {
	preparedStatementPtr, prepareError := databasePtr.Prepare("INSERT INTO " + DUTU.TableName + " VALUES(?, ?)")
	if prepareError != nil {
		return Status{
			HttpStatusCode: http.StatusInternalServerError,
			ErrorMessage:   errorPrepareInsertUserToDatabaseUsersTable + prepareError.Error()}
	}
	_, insertError := preparedStatementPtr.Exec(user.UserName, user.Password)
	if insertError != nil {
		return Status{
			HttpStatusCode: http.StatusInternalServerError,
			ErrorMessage:   errorInsertUserToDatabaseUsersTable + insertError.Error()}
	}
	return Status{
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
	selectRowsPtr, selectError := databasePtr.Query("SELECT * FROM "+DUTU.TableName+" WHERE "+DUTU.UserNameColumnName+" = ?", userName)
	if selectError != nil {
		return dumpUser, Status{
			HttpStatusCode: http.StatusInternalServerError,
			ErrorMessage:   errorSelectGetUserFromDatabaseUsersTable + selectError.Error()}
	}
	defer selectRowsPtr.Close()
	users, getStatus := getAllUsers(selectRowsPtr)
	if getStatus.HttpStatusCode != http.StatusOK {
		return dumpUser, getStatus
	}
	if len(users) != 1 {
		return dumpUser, Status{
			HttpStatusCode: http.StatusInternalServerError,
			ErrorMessage:   errorGetManyUsersGetUserFromDatabaseUsersTable + strconv.Itoa(len(users)) + " user(s)."}
	}
	return users[0], Status{
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
	preparedStatementPtr, prepareError := databasePtr.Prepare("UPDATE " + DUTU.TableName + " SET " + DUTU.PasswordColumnName + " = ? WHERE " + DUTU.UserNameColumnName + " = ?")
	if prepareError != nil {
		return Status{
			HttpStatusCode: http.StatusInternalServerError,
			ErrorMessage:   errorPrepareUpdateUserPasswordToDatabaseUsersTable + prepareError.Error()}
	}
	_, updateError := preparedStatementPtr.Exec(userOfNewPassword.Password, userOfNewPassword.UserName)
	if updateError != nil {
		return Status{
			HttpStatusCode: http.StatusInternalServerError,
			ErrorMessage:   errorUpdateUserPasswordToDatabaseUsersTable + updateError.Error()}
	}
	return Status{
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
	preparedStatementPtr, prepareError := databasePtr.Prepare("DELETE FROM " + DUTU.TableName + " WHERE " + DUTU.UserNameColumnName + " = ?")
	if prepareError != nil {
		return Status{
			HttpStatusCode: http.StatusInternalServerError,
			ErrorMessage:   errorPrepareDeleteUserFromDatabaseUsersTable + prepareError.Error()}
	}
	_, deleteError := preparedStatementPtr.Exec(userName)
	if deleteError != nil {
		return Status{
			HttpStatusCode: http.StatusInternalServerError,
			ErrorMessage:   errorDeleteUserFromDatabaseUsersTable + prepareError.Error()}
	}
	return Status{
		HttpStatusCode: http.StatusOK,
		ErrorMessage:   noError}
}
