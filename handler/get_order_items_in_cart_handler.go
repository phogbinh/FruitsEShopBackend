package handler

import (
	"backend/database"
	"strconv"

	"github.com/gin-gonic/gin"
)

/*
GetOrderItemsInCartHandler is a function for gin to handle GetOrderItemsInCart api
*/
func GetOrderItemsInCartHandler(c *gin.Context) {
	cartID, err := strconv.Atoi(c.Query("cart_id"))
	if err != nil {
		c.Status(400)
	}

	code, items := database.GetOrderItemsInCart(cartID, database.SqlDb)

	c.JSON(code, gin.H{
		"items": items,
	})
}
