package model

type DiscountPolicy struct {
	Code                      string     `json:"code"						binding:"required"`
	Name                      string     `json:"name"						binding:"required"`
	Description               string     `json:"description"				binding:"required"`
	Type                      string     `json:"type"						binding:"required"`
	StaffUserName             string     `json:"staffUserName"				binding:"required"`
	ShippingMinimumOrderPrice NullString `json:"shippingMinimumOrderPrice"	binding:"required"`
	SeasoningsRate            NullString `json:"seasoningsRate"				binding:"required"`
	SeasoningsBeginDate       NullString `json:"seasoningsBeginDate"		binding:"required"`
	SeasoningsEndDate         NullString `json:"seasoningsEndDate"			binding:"required"`
	SpecialEventRate          NullString `json:"specialEventRate"			binding:"required"`
	SpecialEventBeginDate     NullString `json:"specialEventBeginDate"		binding:"required"`
	SpecialEventEndDate       NullString `json:"specialEventEndDate"		binding:"required"`
}
