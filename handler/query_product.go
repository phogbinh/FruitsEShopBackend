package handler

import (
	"backend/database"
	
	"github.com/gin-gonic/gin"
)

/*
QueryProductHandler is a function for gin to handle QueryProduct api
*/
func QueryProductHandler(c *gin.Context) {
	var productName string;
	var staffName string;

	productName = c.Query(database.ProductNameColumnName)
	staffName = c.Query(database.ProductStaffUserNameColumnName)

	code, items := database.QueryProduct(productName, staffName, database.SqlDb)

	c.JSON(code, gin.H{
		"items": items,
	})
}