package handler

import (
	"backend/database"
	. "backend/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

/*
GetOrderItemsInCartHandler is a function for gin to handle GetOrderItemsInCart api
*/
func GetOrderItemsInCartHandler(c *gin.Context) {
	var orderItem OrderItem

	cartID, err := strconv.Atoi(c.Query(database.OrderItemCartIdColumnName))
	if err != nil {
		c.Status(400)
	} else {
		orderItem.CartID = cartID
	}

	code, items := database.GetOrderItemsInCart(&orderItem, database.SqlDb)

	c.JSON(code, gin.H{
		"items": items,
	})
}
