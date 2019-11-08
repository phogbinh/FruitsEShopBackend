package handler

import (
	. "backend/model"
	"fmt"

	"github.com/gin-gonic/gin"
)

/*
SignUpHandler is a function for gin to handle sign up api
*/
func SignUpHandler(c *gin.Context) {
	mail := c.PostForm("mail")
	password := c.PostForm("password")
	userName := c.PostForm("userName")

	user := User{Mail: mail, Password: password, UserName: userName}

	// TODO: store information into db, and checkout the attributes of user
	fmt.Printf("%+v", user)

	c.Status(201)
}
