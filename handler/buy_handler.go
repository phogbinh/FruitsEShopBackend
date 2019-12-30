package handler

import (
	"backend/database"
	. "backend/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func BuyHandler(c *gin.Context) {
	var orderItem OrderItem

	cartID, err := strconv.Atoi(c.Query(database.OrderItemCartIdColumnName))
	if err != nil {
		c.Status(400)
	} else {
		orderItem.CartID = cartID
	}

	code := database.TransactionFromCart(&orderItem, database.SqlDb)

	c.Status(code)
}
