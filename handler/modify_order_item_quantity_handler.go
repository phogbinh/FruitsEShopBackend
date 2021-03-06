package handler

import (
	"backend/database"
	. "backend/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

/*
ModifyOrderItemQuantityHandler is a function for gin to handle ModifyOrderItemQuantity api
*/
func ModifyOrderItemQuantityHandler(c *gin.Context) {
	var orderItem OrderItem

	productID, err := strconv.Atoi(c.Query(database.OrderItemProductIdColumnName))
	if err != nil {
		c.Status(400)
	} else {
		orderItem.ProductID = productID
	}

	cartID, err := strconv.Atoi(c.Query(database.OrderItemCartIdColumnName))
	if err != nil {
		c.Status(400)
	} else {
		orderItem.CartID = cartID
	}

	quantity, err := strconv.Atoi(c.Query(database.OrderItemQuantity))
	if err != nil {
		c.Status(400)
	} else {
		orderItem.Quantity = quantity
	}

	code := database.ModifyOrderItemQuantity(&orderItem, database.SqlDb)

	c.Status(code)
}
