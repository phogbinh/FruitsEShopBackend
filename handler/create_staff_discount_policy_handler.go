package handler

import (
	"database/sql"
	"net/http"

	discountPoliciesTable "backend/database_discount_policies_tables_util/database_discount_policies_table_util"
	discountPolicyTypesTable "backend/database_discount_policies_tables_util/database_discount_policy_types_table_util"
	productsTable "backend/database_products_table_util"
	. "backend/model"
	"backend/util"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	_ "github.com/go-sql-driver/mysql"
)

// CreateStaffDiscountPolicyHandler creates a discount policy for the given staff and responds the discount policy's information.
func CreateStaffDiscountPolicyHandler(databasePtr *sql.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		discountPolicy, getStatus := getDiscountPolicyFromRequest(context)
		if !util.IsStatusOK(getStatus) {
			context.JSON(getStatus.HttpStatusCode, gin.H{util.JsonError: getStatus.ErrorMessage})
			return
		}
		if !discountPolicyTypesTable.IsValidDiscountPolicyType(discountPolicy.Type) {
			context.JSON(http.StatusBadRequest, gin.H{util.JsonError: "The given discount policy type is invalid."})
			return
		}
		insertStatus := discountPoliciesTable.InsertDiscountPolicyToSuperclassAndSubclassTables(discountPolicy, databasePtr)
		if !util.IsStatusOK(insertStatus) {
			context.JSON(insertStatus.HttpStatusCode, gin.H{util.JsonError: insertStatus.ErrorMessage})
			return
		}
		updateStatus := updateProductsTableIfDiscountPolicyIsOfTypeSpecialEvent(discountPolicy, context, databasePtr)
		if !util.IsStatusOK(updateStatus) {
			context.JSON(updateStatus.HttpStatusCode, gin.H{util.JsonError: updateStatus.ErrorMessage})
			return
		}
		discountPolicy, getStatus = discountPoliciesTable.GetDiscountPolicyByCode(discountPolicy.Code, databasePtr)
		if !util.IsStatusOK(getStatus) {
			context.JSON(getStatus.HttpStatusCode, gin.H{util.JsonError: getStatus.ErrorMessage})
			return
		}
		context.JSON(http.StatusOK, discountPolicy)
	}
}

func getDiscountPolicyFromRequest(context *gin.Context) (DiscountPolicy, Status) {
	var discountPolicy DiscountPolicy
	bindError := context.ShouldBindBodyWith(&discountPolicy, binding.JSON) // ShouldBindBodyWith is used here because we want the context data to be reusable by getSpecialEventProductIdsFromRequest.
	if bindError != nil {
		return discountPolicy, util.StatusBadRequest(getDiscountPolicyFromRequest, bindError)
	}
	return discountPolicy, util.StatusOK()
}

func updateProductsTableIfDiscountPolicyIsOfTypeSpecialEvent(discountPolicy DiscountPolicy, context *gin.Context, databasePtr *sql.DB) Status {
	if discountPolicy.Type != discountPolicyTypesTable.TypeSpecialEvent {
		return util.StatusOK()
	}
	return updateSpecialEventProducts(discountPolicy.Code, context, databasePtr)
}

func updateSpecialEventProducts(discountPolicyCode string, context *gin.Context, databasePtr *sql.DB) Status {
	productIds, getStatus := getSpecialEventProductIdsFromRequest(context)
	if !util.IsStatusOK(getStatus) {
		return getStatus
	}
	for _, productId := range productIds.Value {
		updateStatus := productsTable.UpdateProductSpecialEventDiscountPolicyCode(productId, discountPolicyCode, databasePtr)
		if !util.IsStatusOK(updateStatus) {
			return updateStatus
		}
	}
	return util.StatusOK()
}

func getSpecialEventProductIdsFromRequest(context *gin.Context) (SpecialEventProductIds, Status) {
	var productIds SpecialEventProductIds
	bindError := context.ShouldBindBodyWith(&productIds, binding.JSON) // ShouldBindBodyWith is used here because the code does not work with ShouldBindJSON. The cause is unknown.
	if bindError != nil {
		return productIds, util.StatusBadRequest(getSpecialEventProductIdsFromRequest, bindError)
	}
	return productIds, util.StatusOK()
}
