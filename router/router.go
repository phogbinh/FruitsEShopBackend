package router

import (
	"backend/handler"
	"backend/middleware"
	"database/sql"
	"log"

	"backend/symbolutil"

	DUTU "backend/database_users_table_util"
	"github.com/gin-gonic/gin"
)

/*
Register is a place to register rotes
*/
func Register(router *gin.Engine, databasePtr *sql.DB) {
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
	initializeRouterDatabaseUsersTableHandlers(router, databasePtr)

	auth := router.Group("/auth")
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		// TODO: authed api will be here
	}
}

func initializeRouterDatabaseUsersTableHandlers(router *gin.Engine, databasePtr *sql.DB) {
	const userNamePath = ":name"
	router.GET(
		symbolutil.RightSlash+DUTU.TableName,
		handler.ResponseJsonOfAllUsersFromDatabaseUsersTableHandler(databasePtr))

	router.POST(
		symbolutil.RightSlash+DUTU.TableName,
		handler.CreateUserToDatabaseUsersTableAndResponseJsonOfUserHandler(databasePtr))

	router.GET(
		symbolutil.RightSlash+DUTU.TableName+symbolutil.RightSlash+userNamePath,
		handler.ResponseJsonOfUserFromDatabaseUsersTableHandler(databasePtr))

	router.PUT(
		symbolutil.RightSlash+DUTU.TableName+symbolutil.RightSlash+userNamePath,
		handler.UpdateUserPasswordInDatabaseUsersTableAndResponseJsonOfUserHandler(databasePtr))

	router.DELETE(
		symbolutil.RightSlash+DUTU.TableName+symbolutil.RightSlash+userNamePath,
		handler.DeleteUserFromDatabaseUsersTableAndResponseJsonOfUserNameHandler(databasePtr))
}
