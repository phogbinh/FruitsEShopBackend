package handler

import (
	"backend/database"
	"strconv"

	"github.com/gin-gonic/gin"
)

/*
ModifyOrderItemQuantityHandler is a function for gin to handle ModifyOrderItemQuantity api
*/
func ModifyOrderItemQuantityHandler(c *gin.Context) {
	productID, _ := strconv.Atoi(c.Query("p_id"))
	cartID, _ := strconv.Atoi(c.Query("cart_id"))
	quantity, _ := strconv.Atoi(c.Query("quantity"))

	code := database.ModifyOrderItemQuantity(productID, cartID, quantity, database.SqlDb)

	c.Status(code)
}
