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

// SignUpHandler creates an user and responds the user's information.
func SignUpHandler(databasePtr *sql.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		user, getStatus := getUserFromRequest(context)
		if !util.IsStatusOK(getStatus) {
			context.JSON(getStatus.HttpStatusCode, gin.H{util.JsonError: getStatus.ErrorMessage})
			return
		}
		insertStatus := DUTU.InsertUser(user, databasePtr)
		if !util.IsStatusOK(insertStatus) {
			context.JSON(insertStatus.HttpStatusCode, gin.H{util.JsonError: insertStatus.ErrorMessage})
			return
		}
		getUser, getStatus := DUTU.GetUserByUserName(user.UserName, databasePtr)
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
