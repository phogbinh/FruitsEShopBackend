package handler

import (
	"database/sql"
	"net/http"

	discountPoliciesTable "backend/database_discount_policies_tables_util/database_discount_policies_table_util"
	DUTU "backend/database_users_table_util"
	"backend/util"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// GetStaffDiscountPoliciesHandler all discount policies' information of the given staff.
func GetStaffDiscountPoliciesHandler(databasePtr *sql.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		staffUserName := context.Param(DUTU.UserNameColumnName)
		discountPolicies, status := discountPoliciesTable.GetStaffDiscountPolicies(staffUserName, databasePtr)
		if !util.IsStatusOK(status) {
			context.JSON(status.HttpStatusCode, gin.H{util.JsonError: status.ErrorMessage})
			return
		}
		context.JSON(http.StatusOK, discountPolicies)
	}
}
