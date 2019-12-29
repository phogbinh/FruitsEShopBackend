package router

import (
	"database/sql"
	"log"

	discountPoliciesTablesConst "backend/database_discount_policies_tables_util/database_discount_policies_tables_const"
	DUTU "backend/database_users_table_util"
	"backend/handler"
	"backend/middleware"
	"backend/util"

	"github.com/gin-gonic/gin"
)

const (
	userNamePath           = ":" + DUTU.UserNameColumnName
	discountPoliciesPath   = "discount-policies"
	discountPolicyCodePath = ":" + discountPoliciesTablesConst.DiscountPoliciesCodeColumnName
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

	api := router.Group("/api")

	api.POST("/addorderitemtocart", handler.AddOrderItemToCartHandler)
	api.DELETE("/deleteorderitemincart", handler.DeleteOrderItemToCartHandler)
	api.GET("/getorderitemsincart", handler.GetOrderItemsInCartHandler)
	api.PUT("/modifyorderitemquantity", handler.ModifyOrderItemQuantityHandler)
	api.GET("/getcartidwithusername", handler.GetCartIdWithUserNameHandler)

	api.GET("/getstafforder", handler.GetStaffOrderHandler)

	api.GET("/buy", handler.BuyHandler)
	api.GET("/getorder", handler.GetOrderHandler)

	initializeRouterManageUserHandlers(api, databasePtr)

	api.POST("/login", handler.LoginHandler(databasePtr))
	api.POST("/sign-up", handler.SignUpHandler(databasePtr))
	api.GET(util.RightSlash+discountPoliciesPath+util.RightSlash+discountPolicyCodePath, handler.GetDiscountPolicyHandler(databasePtr))
	api.POST("/addproduct", handler.AddProductHandler)
	api.DELETE("/deleteproduct", handler.DeleteProductHandler)
	api.PUT("/modifyproduct", handler.ModifyProductHandler)
	api.GET("/queryproduct", handler.QueryProductHandler)

	auth := router.Group("/api/auth")

	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.PUT(
			util.RightSlash+DUTU.TableName+util.RightSlash+userNamePath,
			handler.UpdateUserPasswordHandler(databasePtr))

		auth.PUT(
			util.RightSlash+DUTU.TableName+util.RightSlash+userNamePath+"/register-staff",
			handler.RegisterStaffHandler(databasePtr))

		auth.POST(
			util.RightSlash+userNamePath+util.RightSlash+discountPoliciesPath,
			handler.CreateStaffDiscountPolicyHandler(databasePtr))

		auth.GET(
			util.RightSlash+userNamePath+util.RightSlash+discountPoliciesPath,
			handler.GetStaffDiscountPoliciesHandler(databasePtr))

		auth.DELETE(
			util.RightSlash+userNamePath+util.RightSlash+discountPoliciesPath+util.RightSlash+discountPolicyCodePath,
			handler.DeleteStaffDiscountPolicyHandler(databasePtr))
	}
}

func initializeRouterManageUserHandlers(router *gin.RouterGroup, databasePtr *sql.DB) {
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
