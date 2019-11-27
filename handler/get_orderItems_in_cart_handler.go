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
	cartID, _ := strconv.Atoi(c.Query("cart_id"))

	code, items := database.GetOrderItemsInCart(cartID, database.SqlDb)

	c.JSON(code, gin.H{
		"items": items,
	})
}
