package handler

import (
	"backend/database"
	"strconv"

	"github.com/gin-gonic/gin"
)

/*
AddOrderItemToCartHandler is a function for gin to handle AddOrderItemToCart api
*/
func AddOrderItemToCartHandler(c *gin.Context) {
	productID, _ := strconv.Atoi(c.Query("p_id"))
	cartID, _ := strconv.Atoi(c.Query("cart_id"))
	quantity, _ := strconv.Atoi(c.Query("quantity"))

	code := database.AddOrderItemToCart(productID, cartID, quantity, database.SqlDb)

	c.Status(code)
}
