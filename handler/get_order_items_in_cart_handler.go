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
	var addToCart AddToCart

	cartID, err := strconv.Atoi(c.Query("cart_id"))
	if err != nil {
		c.Status(400)
	} else {
		addToCart.CartID = cartID
	}

	code, items := database.GetOrderItemsInCart(&addToCart, database.SqlDb)

	c.JSON(code, gin.H{
		"items": items,
	})
}
