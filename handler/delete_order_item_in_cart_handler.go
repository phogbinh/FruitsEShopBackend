package handler

import (
	"backend/database"
	"strconv"

	"github.com/gin-gonic/gin"
)

/*
DeleteOrderItemToCartHandler is a function for gin to handle DeleteOrderItemToCart api
*/
func DeleteOrderItemToCartHandler(c *gin.Context) {
	productID, err := strconv.Atoi(c.Query("p_id"))
	if err != nil {
		c.Status(400)
	}

	cartID, err := strconv.Atoi(c.Query("cart_id"))
	if err != nil {
		c.Status(400)
	}

	code := database.DeleteOrderItemInCart(productID, cartID, database.SqlDb)

	c.Status(code)
}
