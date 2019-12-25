package model

type DiscountPolicy struct {
	Code                      string `json:"code"						binding:"required"`
	Name                      string `json:"name"						binding:"required"`
	Description               string `json:"description"				binding:"required"`
	Type                      string `json:"type"						binding:"required"`
	StaffUserName             string `json:"staffUserName"				binding:"required"`
	ShippingMinimumOrderPrice string `json:"shippingMinimumOrderPrice"	binding:"required"`
	SeasoningsRate            string `json:"seasoningsRate"				binding:"required"`
	SeasoningsBeginDate       string `json:"seasoningsBeginDate"		binding:"required"`
	SeasoningsEndDate         string `json:"seasoningsEndDate"			binding:"required"`
	SpecialEventRate          string `json:"specialEventRate"			binding:"required"`
	SpecialEventBeginDate     string `json:"specialEventBeginDate"		binding:"required"`
	SpecialEventEndDate       string `json:"specialEventEndDate"		binding:"required"`
}
