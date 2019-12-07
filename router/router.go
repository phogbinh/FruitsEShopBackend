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

const (
	authorizationPath = "auth"
	userNamePath      = ":" + DUTU.UserNameColumnName
)

/*
Register is a place to register rotes
*/
func Register(router *gin.Engine, databasePtr *sql.DB) {
	authMiddleware, err := middleware.NewAuthMiddleware()
	if err != nil {
		log.Panicln(err)
	}

	router.POST("/addorderitemtocart", handler.AddOrderItemToCartHandler)
	router.DELETE("/deleteorderitemincart", handler.DeleteOrderItemToCartHandler)
	router.GET("/getorderitemsincart", handler.GetOrderItemsInCartHandler)
	router.PUT("/modifyorderitemquantity", handler.ModifyOrderItemQuantityHandler)
	router.POST("/login", handler.LoginHandler(databasePtr))
	initializeRouterDatabaseUsersTableHandlers(router, databasePtr)

	auth := router.Group(util.RightSlash + authorizationPath)
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		// TODO: authed api will be here
	}
}

func initializeRouterDatabaseUsersTableHandlers(router *gin.Engine, databasePtr *sql.DB) {
	router.POST(
		util.RightSlash+DUTU.TableName,
		handler.SignUpHandler(databasePtr))

	router.GET(
		util.RightSlash+DUTU.TableName,
		handler.RespondJsonOfAllUsersFromDatabaseUsersTableHandler(databasePtr))

	router.GET(
		util.RightSlash+DUTU.TableName+util.RightSlash+userNamePath,
		handler.RespondJsonOfUserByUserNameFromDatabaseUsersTableHandler(databasePtr))

	router.GET(util.RightSlash+"user",
		handler.RespondJsonOfUserByMailFromDatabaseUsersTableHandler(databasePtr))

	router.PUT(
		util.RightSlash+DUTU.TableName+util.RightSlash+userNamePath,
		handler.UpdateUserPasswordHandler(databasePtr))

	router.DELETE(
		util.RightSlash+DUTU.TableName+util.RightSlash+userNamePath,
		handler.DeleteUserFromDatabaseUsersTable(databasePtr))
}
