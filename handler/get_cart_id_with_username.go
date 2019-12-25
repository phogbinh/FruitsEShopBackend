package handler

import (
	"backend/database"
	DUTU "backend/database_users_table_util"
	. "backend/model"

	"github.com/gin-gonic/gin"
)

/*
GetCartIdWithUserNameHandler is a function for gin to handle GetCartIdWithUserName api
*/
func GetCartIdWithUserName(c *gin.Context) {
	var user User

	userName := c.Param(DUTU.UserNameColumnName)
	user.UserName = userName

	code, cartId := database.GetCartIdWithUsername(&user, database.SqlDb)

	c.JSON(code, gin.H{
		"CartId": cartId,
	})
}
