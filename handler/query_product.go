package handler

import (
	"backend/database"
	"strconv"
	"../model/product.go"
	
	"github.com/gin-gonic/gin"
)

/*
QueryProductHandler is a function for gin to handle QueryProduct api
*/
func QueryProductHandler(c *gin.Context) {
	string productName;

	productName, _ := c.Query("p_name")
	staffName, _ := c.Query("s_username")

	status, items := database.QueryProduct(productName, staffName, database.SqlDb)

	c.JSON(code, gin.H{
		"items": items,
	})
}