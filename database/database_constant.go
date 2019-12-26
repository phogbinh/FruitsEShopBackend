package database

import "backend/database_products_table_util"

const (
	AdminTableName          = "admin"
	AdminUserNameColumnName = "AdminUserName"
	AdminPasswordColumnName = "Password"
	AdminNicknameColumnName = "Nickname"
	AdminEmailColumnName    = "Mail"
)

const (
	ProductTableName               = "product"
	ProductIdColumnName            = database_products_table_util.IdColumnName
	ProductStaffUserNameColumnName = "StaffUserName"
	ProductDescriptionColumnName   = "Description"
	ProductNameColumnName          = "Pname"
	ProductCategoryColumnName      = "Category"
	ProductSourceColumnName        = "Source"
	ProductPriceColumnName         = "Price"
	ProductInventoryColumnName     = "Inventory"
	ProductSoldQuantityColumnName  = "SoldQuantity"
	ProductOnSaleDateColumnName    = "OnSaleDate"
	productDiscountPolicyCodeColumnName = "SpecialEventDiscountPolicyCode"
)

const (
	QA_TableName                  = "qa"
	QA_ProductIdColumnName        = "ProductId"
	QA_StaffUserNameColumnName    = "StaffUserName"
	QA_CustomerUserNameColumnName = "CustomerUserName"
	QA_QuestionColumnname         = "Question"
	QA_AnswerColumnName           = "Answer"
	QA_AskDatetimeColumnName      = "AskDatetime"
	QA_AnsDatetimeColumnName      = "AnsDatetime"
)

const (
	CartTableName    = "cart"
	CartIdColumnName = "CartId"
)

const (
	ActivityTableName                = "activity"
	ActivityIdColumnName             = "ActivityId"
	ActivityNameColumnName           = "Name"
	ActivityStartDateColumnName      = "StartDate"
	ActivityEndDateColumnName        = "EndDate"
	ActivityLowestDiscountColumnName = "LowestDiscount"
)

const (
	OrderItemTableName           = "order_item"
	OrderItemCartIdColumnName    = "CartId"
	OrderItemProductIdColumnName = "ProductId"
	OrderItemQuantity            = "Quantity"
)

const (
	CustomerOwnCartTableName                  = "customer_own_cart"
	CustomerOwnCartCartIdColumnName           = "CartId"
	CustomerOwnCartCustomerUserNameColumnName = "CustomerUserName"
)

const (
	TradeTableName                 = "trade"
	TradeCartIdColumnName          = "CartId"
	TradeProductIdColumnName       = "ProductId"
	TradeProductQuantityColumnName = "Quantity"
	TradeDateTimeColumnName        = "DateTime"
)

const (
	CustomerEvaluateTableName                  = "c_evaluate"
	CustomerEvaluateCustomerUserNameColumnName = "CustomerUserName"
	CustomerEvaluateStaffUserNameColumnName    = "StaffUserName"
	CustomerEvaluateFeedbackColumnName         = "Feedback"
)

const (
	StaffEvaluateTableName                  = "s_evaluate"
	StaffEvaluateCustomerUserNameColumnName = "CustomerUserName"
	StaffEvaluateStaffUserNameColumnName    = "StaffUserName"
	StaffEvaluateFeedbackColumnName         = "Feedback"
)

const (
	ProductEvaluateTableName                  = "p_evaluate"
	ProductEvaluateProductIdColumnName        = "ProductId"
	ProductEvaluateCustomerUserNameColumnName = "CustomerUserName"
	ProductEvaluateFeedbackColumnName         = "Feedback"
)

const (
	ManageTableName               = "manage"
	ManageAdminUserNameColumnName = "AdminUserName"
	ManageUserNameColumnName      = "UserName"
)

const (
	HoldTableName               = "hold"
	HoldAdminUserNameColumnName = "AdminUserName"
	HoldActivityIdColumnName    = "ActivityId"
)

const (
	StaffJoinActivityTableName               = "s_join"
	StaffJoinActivityStaffUserNameColumnName = "StaffUserName"
	StaffJoinActivityActivityIdColumnName    = "ActivityId"
)

const (
	ProductJoinActivityTableName            = "p_join"
	ProductJoinActivityProductIdColumnName  = "ProductId"
	ProductJoinActivityActivityIdColumnName = "ActivityId"
	ProductJoinActivityQuantityColumnName   = "Quantity"
	ProductJoinActivityDiscountColumnName   = "Discount"
)

const (
	TakeOffTableName           = "take_off"
	TakeOffUserNameColumnName  = "UserName"
	TakeOffProductIdColumnName = "ProductId"
)
