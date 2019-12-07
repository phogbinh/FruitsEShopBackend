package handler

import (
	"database/sql"
	"log"
	"net/http"

	DUTU "backend/database_users_table_util"
	"backend/middleware"
	. "backend/model"
	"backend/util"

	"github.com/gin-gonic/gin"
)

/*
LoginHandler is a function for gin to handle login api
*/
func LoginHandler(databasePtr *sql.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		authMiddleware, err := middleware.NewAuthMiddleware()
		if err != nil {
			log.Printf("JWT Error:" + err.Error())
			context.Status(http.StatusInternalServerError)
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
			context.Status(http.StatusBadRequest)
			return
		}
		context.Status(http.StatusOK)
		authMiddleware.LoginHandler(context)
	}
}

func getLoginFromRequest(context *gin.Context) (Login, Status) {
	var login Login
	bindError := context.ShouldBindJSON(&login)
	if bindError != nil {
		return login, util.StatusBadRequest(getLoginFromRequest, bindError)
	}
	return login, util.StatusOK()
}
