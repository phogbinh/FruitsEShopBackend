package handler

import (
	"backend/database"
	"strconv"
	"../model/product.go"

	"github.com/gin-gonic/gin"
)

/*
AddProductHandler is a function for gin to handle AddProduct api
*/
func AddProductHandler(c *gin.Context) {
	var productInfo Product
	
	productInfo.StaffName, _ := c.Query("s_username")
	productInfo.Description, _ := c.Query("description")
	productInfo.Name, _ := c.Query("p_name")
	productInfo.Category, _ := c.Query("category")
	productInfo.Source, _ := c.Query("source")
	productInfo.Price, _ := strconv.Atoi(c.Query("price"))
	productInfo.Inventory, _ := strconv.Atoi(c.Query("inventory"))
	productInfo.Quantity, _ := strconv.Atoi(c.Query("sold_quantity"))
	productInfo.SaleDate, _ := c.Query("onsale_date")

	status := database.AddProduct(&productInfo, database.SqlDb)

	c.Status(status)
}
