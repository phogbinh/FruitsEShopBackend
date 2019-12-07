package handler

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	DUTU "backend/database_users_table_util"
	"backend/middleware"

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

		fmt.Printf("mail is" + mail + "password is" + password)

		context.Status(http.StatusOK)

		authMiddleware.LoginHandler(context)
	}
}
