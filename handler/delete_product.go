package handler

import (
	"backend/database"
	"strconv"

	"github.com/gin-gonic/gin"
)

/*
DeleteProductHandler is a function for gin to handle DeleteProduct api
*/
func DeleteProductHandler(c *gin.Context) {
	productID, _ := strconv.Atoi(c.Query("p_id"))

	status := database.DeleteProduct(productID)

	c.Status(status)
}
