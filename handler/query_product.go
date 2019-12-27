package handler

import (
	"backend/database"
	"strconv"

	"github.com/gin-gonic/gin"
)

/*
QueryProductHandler is a function for gin to handle QueryProduct api
*/
func QueryProductHandler(c *gin.Context) {
	var productName string;
	var staffName string;
	var pid int;

	productName = c.Query(database.ProductNameColumnName)
	staffName = c.Query(database.ProductStaffUserNameColumnName)
	productID, err := strconv.Atoi(c.Query(database.ProductIdColumnName))
	if err != nil {
		c.Status(400)
	} else {
		pid = productID
	}

	code, items := database.QueryProduct(productName, staffName, pid, database.SqlDb)

	c.JSON(code, gin.H{
		"items": items,
	})
}