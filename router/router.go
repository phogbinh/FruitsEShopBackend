package router

import (
	"backend/handler"
	"backend/middleware"
	"log"

	"github.com/gin-gonic/gin"
)

/*
Register is a place to register rotes
*/
func Register(router *gin.Engine) {
	authMiddleware, err := middleware.NewAuthMiddleware()
	if err != nil {
		log.Panicln(err)
	}

	router.POST("/login", handler.LoginHandler)
	router.POST("/signup", handler.SignUpHandler)
	router.POST("/addorderitemtocart", handler.AddOrderItemToCartHandler)
	router.DELETE("/deleteorderitemincart", handler.DeleteOrderItemToCartHandler)
	router.GET("/getorderitemsincart", handler.GetOrderItemsInCartHandler)
	router.PUT("/modifyorderitemquantity", handler.ModifyOrderItemQuantityHandler)

	auth := router.Group("/auth")
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		// TODO: authed api will be here
	}
}
