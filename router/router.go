package router

import (
	"database/sql"
	"log"

	DUTU "backend/database_users_table_util"
	"backend/handler"
	"backend/middleware"
	"backend/util"

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
	const userNamePath = ":" + DUTU.UserNameColumnName
	router.GET(
		util.RightSlash+DUTU.TableName,
		handler.RespondJsonOfAllUsersFromDatabaseUsersTableHandler(databasePtr))

	router.POST(
		util.RightSlash+DUTU.TableName,
		handler.CreateUserToDatabaseUsersTableAndRespondJsonOfUserHandler(databasePtr))

	router.GET(
		util.RightSlash+DUTU.TableName+util.RightSlash+userNamePath,
		handler.RespondJsonOfUserFromDatabaseUsersTableHandler(databasePtr))

	router.PUT(
		util.RightSlash+DUTU.TableName+util.RightSlash+userNamePath,
		handler.UpdateUserPasswordInDatabaseUsersTableAndRespondJsonOfUserHandler(databasePtr))

	router.DELETE(
		util.RightSlash+DUTU.TableName+util.RightSlash+userNamePath,
		handler.DeleteUserFromDatabaseUsersTableAndRespondJsonOfUserNameHandler(databasePtr))
}
