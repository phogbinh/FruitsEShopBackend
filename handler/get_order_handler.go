package handler

import (
	"backend/database"
	. "backend/model"

	"github.com/gin-gonic/gin"
)

/// GetOrderHandler is a function for gin to handle get order
func GetOrderHandler(c *gin.Context) {
	var user User

	userName := c.Query("username")
	user.UserName = userName

	code, message := database.GetOrder(user, database.SqlDb)

	c.JSON(code, gin.H{
		"items": message,
	})

}
