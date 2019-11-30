package model

type User struct {
	Mail         string `json:"mail"			binding:"required"`
	Password     string `json:"password"		binding:"required"`
	UserName     string `json:"userName"		binding:"required"`
	Nickname     string `json:"nickname"		binding:"required"`
	Fname        string `json:"fname"			binding:"required"`
	Lname        string `json:"lname"			binding:"required"`
	Phone        string `json:"phone"			binding:"required"`
	Location     string `json:"location"		binding:"required"`
	Money        string `json:"money"			binding:"required"`
	Introduction string `json:"introduction"	binding:"required"`
}
