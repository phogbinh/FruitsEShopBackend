package handler

import (
	"backend/database"
	"strconv"

	"github.com/gin-gonic/gin"
)

/// GetOrderHandler is a function for gin to handle get order
func GetOrderHandler(c *gin.Context) {

	cartID, err := strconv.Atoi(c.Query(database.OrderItemCartIdColumnName))
	if err != nil {
		c.Status(400)
	}

	code, message := database.GetOrder(cartID, database.SqlDb)

	c.JSON(code, gin.H{
		"items": message,
	})

}
