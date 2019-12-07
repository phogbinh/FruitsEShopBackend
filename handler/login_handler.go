package handler

import (
	"database/sql"
	"log"
	"net/http"

	DUTU "backend/database_users_table_util"
	"backend/middleware"
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
		mail := context.PostForm(DUTU.MailColumnName)
		password := context.PostForm(DUTU.PasswordColumnName)
		user, status := DUTU.GetUserByMail(mail, databasePtr)
		if !util.IsStatusOK(status) {
			context.JSON(status.HttpStatusCode, gin.H{util.JsonError: status.ErrorMessage})
			return
		}
		if user.Password != password {
			context.Status(http.StatusBadRequest)
			return
		}
		context.Status(http.StatusOK)
		authMiddleware.LoginHandler(context)
	}
}
