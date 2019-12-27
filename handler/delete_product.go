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
	var pid int
	productID, err := strconv.Atoi(c.Query(database.ProductIdColumnName))
	if err != nil {
		c.Status(400)
	} else {
		pid = productID
	}

	code := database.DeleteProduct(pid, database.SqlDb)

	c.Status(code)
}
