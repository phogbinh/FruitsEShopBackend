package database

const (
	AdminTableName          = "admin"
	AdminUserNameColumnName = "AdminUserName"
	AdminPasswordColumnName = "Password"
	AdminNicknameColumnName = "Nickname"
	AdminEmailColumnName    = "Mail"
)

const (
	ProductTableName               = "product"
	ProductIdColumnName            = "ProductId"
	ProductStaffUserNameColumnName = "StaffUserName"
	ProductDescriptionColumnName   = "Description"
	ProductNameColumnName          = "Pname"
	ProductCategoryColumnName      = "Category"
	ProductSourceColumnName        = "Source"
	ProductPriceColumnName         = "Price"
	ProductInventoryColumnName     = "Inventory"
	ProductSoldQuantityColumnName  = "SoldQuantity"
	ProductOnSaleDataColumnName    = "OnSaleDate"
)

const (
	QA_TableName                  = "qa"
	QA_ProductIdColumnName        = "ProductId"
	QA_StaffUserNameColumnName    = "StaffUserName"
	QA_CustomerUserNameColumnName = "CustomerUserName"
	QA_QuestionColumnname         = "Question"
	QA_AnswerColumnName           = "Answer"
	QA_AskDateColumnName          = "AskDate"
	QA_AskTimeColumnName          = "AskTime"
	QA_AnsDateColumnName          = "AnsDate"
	QA_AnsTimeColumnName          = "AnsTime"
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
	AddToCartTableName           = "add_to_cart"
	AddToCartCartIdColumnName    = "CartId"
	AddToCartProductIdColumnName = "ProductId"
	AddToCartQuantityColumnName  = "Quantity"
)

const (
	BuyTableName                  = "buy"
	BuyCustomerUserNameColumnName = "CustomerUserName"
	BuyCartIdColumnName           = "BuyCartId"
	BuyBuyDateColumnName          = "BuyDate"
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
	TakeOffTableName               = "take_off"
	TakeOffAdminUserNameColumnName = "ActivityUserName"
	TakeOffProductIdColumnName     = "ProductId"
)
