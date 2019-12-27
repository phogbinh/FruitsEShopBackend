package handler

import (
	"backend/database"
	"strconv"
	. "backend/model"
	"github.com/gin-gonic/gin"
)

/*
AddProductHandler is a function for gin to handle AddProduct api
*/
func AddProductHandler(c *gin.Context) {
	var productInfo Product
	
	staffName := c.Query(database.ProductStaffUserNameColumnName)
	productInfo.StaffName = staffName;

	description := c.Query(database.ProductDescriptionColumnName)
	productInfo.Description = description

	name := c.Query(database.ProductNameColumnName)
	productInfo.Name = name

	category := c.Query(database.ProductCategoryColumnName)
	productInfo.Category = category

	source := c.Query(database.ProductSourceColumnName)
	productInfo.Source = source

	price, err := strconv.Atoi(c.Query(database.ProductPriceColumnName))
	if err != nil {
		c.Status(400)
	} else {
		productInfo.Price = price
	}

	inventory, err := strconv.Atoi(c.Query(database.ProductInventoryColumnName))
	if err != nil {
		c.Status(400)
	} else {
		productInfo.Inventory = inventory
	}

	quantity, err := strconv.Atoi(c.Query(database.ProductSoldQuantityColumnName))
	if err != nil {
		c.Status(400)
	} else {
		productInfo.Quantity = quantity
	}

	saledate := c.Query(database.ProductOnSaleDateColumnName)
	productInfo.SaleDate = saledate;
	
	imageSrc:= c.Query(database.ProductImageSourceColumnName)
	productInfo.ImageSrc = imageSrc


	code := database.AddProduct(&productInfo, database.SqlDb)

	c.Status(code)
}
