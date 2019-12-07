package handler

import (
	"database/sql"
	"fmt"
	"log"

	"backend/middleware"

	"github.com/gin-gonic/gin"
)

/*
LoginHandler is a function for gin to handle login api
*/
func LoginHandler(databasePtr *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		authMiddleware, err := middleware.NewAuthMiddleware()
		if err != nil {
			log.Printf("JWT Error:" + err.Error())
			c.Status(500)
		}

		mail := c.PostForm("mail")
		password := c.PostForm("password")

		fmt.Printf("mail is" + mail + "password is" + password)

		c.Status(201)

		authMiddleware.LoginHandler(c)
	}
}
