package handler

import (
	"backend/database"
	"backend/database_users_table_util"

	"github.com/gin-gonic/gin"
)

/*
GetStaffOrder is a function for gin to handle GetOrderItemsInCart api
*/
func GetStaffOrderHandler(c *gin.Context) {
	staffUserName := c.Query(database_users_table_util.UserNameColumnName)

	code, items := database.GetStaffOrder(staffUserName, database.SqlDb)

	c.JSON(code, gin.H{
		"items": items,
	})
}
