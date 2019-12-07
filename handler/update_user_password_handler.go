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

// UpdateUserPasswordHandler updates an user's password and responds the user's information.
func UpdateUserPasswordHandler(databasePtr *sql.DB) gin.HandlerFunc {
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
