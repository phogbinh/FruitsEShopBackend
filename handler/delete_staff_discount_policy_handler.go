package handler

import (
	"database/sql"
	"net/http"

	discountPoliciesTable "backend/database_discount_policies_tables_util/database_discount_policies_table_util"
	discountPoliciesTablesConst "backend/database_discount_policies_tables_util/database_discount_policies_tables_const"
	"backend/util"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// DeleteStaffDiscountPolicyHandler deletes a discount policy of the given staff.
func DeleteStaffDiscountPolicyHandler(databasePtr *sql.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		discountPolicyCode := context.Param(discountPoliciesTablesConst.DiscountPoliciesCodeColumnName)
		deleteStatus := discountPoliciesTable.DeleteDiscountPolicy(discountPolicyCode, databasePtr)
		if !util.IsStatusOK(deleteStatus) {
			context.JSON(deleteStatus.HttpStatusCode, gin.H{util.JsonError: deleteStatus.ErrorMessage})
			return
		}
		isExistingDiscountPolicy, getStatus := discountPoliciesTable.IsExistingDiscountPolicy(discountPolicyCode, databasePtr)
		if !util.IsStatusOK(getStatus) {
			context.JSON(getStatus.HttpStatusCode, gin.H{util.JsonError: getStatus.ErrorMessage})
			return
		}
		if isExistingDiscountPolicy {
			context.JSON(http.StatusInternalServerError, gin.H{util.JsonError: "Discount policy still exists."})
			return
		}
		context.Status(http.StatusOK)
	}
}
