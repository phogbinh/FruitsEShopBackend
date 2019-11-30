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
	productID, err := strconv.Atoi(c.Query("p_id"))
	if err != nil {
		c.Status(400)
	}

	cartID, err := strconv.Atoi(c.Query("cart_id"))
	if err != nil {
		c.Status(400)
	}

	quantity, err := strconv.Atoi(c.Query("quantity"))
	if err != nil {
		c.Status(400)
	}

	code := database.ModifyOrderItemQuantity(productID, cartID, quantity, database.SqlDb)

	c.Status(code)
}
