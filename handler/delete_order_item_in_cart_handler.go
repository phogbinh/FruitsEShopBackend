package handler

import (
	"backend/database"
	. "backend/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

/*
DeleteOrderItemToCartHandler is a function for gin to handle DeleteOrderItemToCart api
*/
func DeleteOrderItemToCartHandler(c *gin.Context) {
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

	code := database.DeleteOrderItemInCart(&orderItem, database.SqlDb)

	c.Status(code)
}
