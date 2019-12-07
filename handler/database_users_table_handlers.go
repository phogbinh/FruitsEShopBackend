package handler

import (
	"database/sql"
	"net/http"

	DUTU "backend/database_users_table_util"
	. "backend/model"
	"backend/util"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type password struct {
	Value string `json:"password" binding:"required"`
}

// RespondJsonOfAllUsersFromDatabaseUsersTableHandler responds all users' information.
func RespondJsonOfAllUsersFromDatabaseUsersTableHandler(databasePtr *sql.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		users, status := DUTU.GetAllUsers(databasePtr)
		if !util.IsStatusOK(status) {
			context.JSON(status.HttpStatusCode, gin.H{util.JsonError: status.ErrorMessage})
			return
		}
		context.JSON(http.StatusOK, users)
	}
}

// RespondJsonOfUserByUserNameFromDatabaseUsersTableHandler responds an user's information.
func RespondJsonOfUserByUserNameFromDatabaseUsersTableHandler(databasePtr *sql.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		userName := context.Param(DUTU.UserNameColumnName)
		user, status := DUTU.GetUserByUserName(userName, databasePtr)
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
		user, status := DUTU.GetUserByMail(mail, databasePtr)
		if !util.IsStatusOK(status) {
			context.JSON(status.HttpStatusCode, gin.H{util.JsonError: status.ErrorMessage})
			return
		}
		context.JSON(http.StatusOK, user)
	}
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
		updateStatus := DUTU.UpdateUserPassword(userName, userNewPassword.Value, databasePtr)
		if !util.IsStatusOK(updateStatus) {
			context.JSON(updateStatus.HttpStatusCode, gin.H{util.JsonError: updateStatus.ErrorMessage})
			return
		}
		getUser, getStatus := DUTU.GetUserByUserName(userName, databasePtr)
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

// DeleteUserFromDatabaseUsersTable deletes an user.
func DeleteUserFromDatabaseUsersTable(databasePtr *sql.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		userName := context.Param(DUTU.UserNameColumnName)
		deleteStatus := DUTU.DeleteUser(userName, databasePtr)
		if !util.IsStatusOK(deleteStatus) {
			context.JSON(deleteStatus.HttpStatusCode, gin.H{util.JsonError: deleteStatus.ErrorMessage})
			return
		}
		isExistingUser, getStatus := DUTU.IsExistingUser(userName, databasePtr)
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
