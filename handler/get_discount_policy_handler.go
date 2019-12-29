package handler

import (
	"database/sql"
	"net/http"

	discountPoliciesTable "backend/database_discount_policies_tables_util/database_discount_policies_table_util"
	discountPoliciesTablesConst "backend/database_discount_policies_tables_util/database_discount_policies_tables_const"
	"backend/util"

	"github.com/gin-gonic/gin"
)

// GetDiscountPolicyHandler responds a discount policy's information by the given code.
func GetDiscountPolicyHandler(databasePtr *sql.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		discountPolicyCode := context.Param(discountPoliciesTablesConst.DiscountPoliciesCodeColumnName)
		discountPolicy, getStatus := discountPoliciesTable.GetDiscountPolicyByCode(discountPolicyCode, databasePtr)
		if !util.IsStatusOK(getStatus) {
			context.JSON(getStatus.HttpStatusCode, gin.H{util.JsonError: getStatus.ErrorMessage})
			return
		}
		context.JSON(http.StatusOK, discountPolicy)
	}
}
