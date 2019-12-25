package handler

import (
	"backend/database"
	DUTU "backend/database_users_table_util"
	"fmt"

	"github.com/gin-gonic/gin"
)

/*
GetCartIdWithUserNameHandler is a function for gin to handle GetCartIdWithUserName api
*/
func GetCartIdWithUserNameHandler(c *gin.Context) {
	userName := c.Query(DUTU.UserNameColumnName)

	fmt.Println("user name = ", userName)

	code, cartId := database.GetCartIdWithUsername(userName, database.SqlDb)

	c.JSON(code, gin.H{
		"CartId": cartId,
	})
}
