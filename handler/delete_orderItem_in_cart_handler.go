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
	productID, _ := strconv.Atoi(c.Query("p_id"))
	cartID, _ := strconv.Atoi(c.Query("cart_id"))

	code := database.DeleteOrderItemInCart(productID, cartID, database.SqlDb)

	c.Status(code)
}
