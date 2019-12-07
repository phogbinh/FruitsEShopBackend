package handler

import (
	"database/sql"
	"net/http"

	DUTU "backend/database_users_table_util"
	"backend/middleware"
	. "backend/model"
	"backend/util"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

/*
LoginHandler is a function for gin to handle login api
*/
func LoginHandler(databasePtr *sql.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		authMiddleware, err := middleware.NewAuthMiddleware()
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{util.JsonError: "JWT Error:" + err.Error()})
			return
		}
		login, getLoginStatus := getLoginFromRequest(context)
		if !util.IsStatusOK(getLoginStatus) {
			context.JSON(getLoginStatus.HttpStatusCode, gin.H{util.JsonError: getLoginStatus.ErrorMessage})
			return
		}
		user, getUserStatus := DUTU.GetUserByMail(login.Mail, databasePtr)
		if !util.IsStatusOK(getUserStatus) {
			context.JSON(getUserStatus.HttpStatusCode, gin.H{util.JsonError: getUserStatus.ErrorMessage})
			return
		}
		if user.Password != login.Password {
			context.JSON(http.StatusBadRequest, gin.H{util.JsonError: "Incorrect password."})
			return
		}
		context.Status(http.StatusOK)
		authMiddleware.LoginHandler(context)
	}
}

func getLoginFromRequest(context *gin.Context) (Login, Status) {
	var login Login
	bindError := context.ShouldBindBodyWith(&login, binding.JSON) // Since unlike ShouldBindWith, ShouldBindBodyWith puts back data into context after reading, it is used here so that context data can be passed down to authMiddleware.LoginHandler(context) for data fetching.
	if bindError != nil {
		return login, util.StatusBadRequest(getLoginFromRequest, bindError)
	}
	return login, util.StatusOK()
}
