package handler

import (
	"backend/database"
	"strconv"
	"../model/product.go"
	
	"github.com/gin-gonic/gin"
)

type test struct {
	attr1     string
	attr2     int
}

/*
ModifyProductHandler is a function for gin to handle ModifyProduct api
*/
func ModifyProductHandler(c *gin.Context) {
	var productInfo Product

	productID, _ := strconv.Atoi(c.Query("p_id"))
	productInfo.StaffName, _ := c.Query("s_username")
	productInfo.Description, _ := c.Query("description")
	productInfo.Name, _ := c.Query("p_name")
	productInfo.Category, _ := c.Query("category")
	productInfo.Source, _ := c.Query("source")
	productInfo.Price, _ := strconv.Atoi(c.Query("price"))
	productInfo.Inventory, _ := strconv.Atoi(c.Query("inventory"))
	productInfo.Quantity, _ := strconv.Atoi(c.Query("sold_quantity"))
	productInfo.SaleDate, _ := c.Query("onsale_date")

	status := database.ModifyProduct(productID, &productInfo, database.SqlDb)

	c.Status(status)
}
