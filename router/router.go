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
	userNamePath = ":" + DUTU.UserNameColumnName
)

/*
Register is a place to register rotes
*/
func Register(router *gin.Engine, databasePtr *sql.DB) {
	authMiddleware, err := middleware.NewAuthMiddleware()
	if err != nil {
		log.Panicln(err)
	}

	router.Use(middleware.NewCORSMiddleware())

	router.POST("/addorderitemtocart", handler.AddOrderItemToCartHandler)
	router.DELETE("/deleteorderitemincart", handler.DeleteOrderItemToCartHandler)
	router.GET("/getorderitemsincart", handler.GetOrderItemsInCartHandler)
	router.PUT("/modifyorderitemquantity", handler.ModifyOrderItemQuantityHandler)
	router.GET("/getcartidwithusername", handler.GetCartIdWithUserName)

	router.GET("/buy", handler.BuyHandler)
	initializeRouterManageUserHandlers(router, databasePtr)

	router.POST("/login", handler.LoginHandler(databasePtr))
	router.POST("/sign-up", handler.SignUpHandler(databasePtr))

	auth := router.Group("/auth")

	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.PUT(
			util.RightSlash+DUTU.TableName+util.RightSlash+userNamePath,
			handler.UpdateUserPasswordHandler(databasePtr))

		auth.PUT(
			util.RightSlash+DUTU.TableName+util.RightSlash+userNamePath+"/register-staff",
			handler.RegisterStaffHandler(databasePtr))
	}
}

func initializeRouterManageUserHandlers(router *gin.Engine, databasePtr *sql.DB) {
	router.GET(
		util.RightSlash+DUTU.TableName,
		handler.RespondJsonOfAllUsersFromDatabaseUsersTableHandler(databasePtr))
	router.GET(
		util.RightSlash+DUTU.TableName+util.RightSlash+userNamePath,
		handler.RespondJsonOfUserByUserNameFromDatabaseUsersTableHandler(databasePtr))
	router.GET("/user",
		handler.RespondJsonOfUserByMailFromDatabaseUsersTableHandler(databasePtr))
	router.DELETE(
		util.RightSlash+DUTU.TableName+util.RightSlash+userNamePath,
		handler.DeleteUserFromDatabaseUsersTable(databasePtr))
}
