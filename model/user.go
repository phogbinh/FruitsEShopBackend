package model

type User struct {
	Mail         string `json:"mail"`
	Password     string `json:"password" binding:"required"`
	UserName     string `json:"userName" binding:"required"`
	Nickname     string `json:"nickname"`
	Fname        string `json:"fname"`
	Lname        string `json:"lname"`
	Phone        string `json:"phone"`
	Location     string `json:"location"`
	Money        string `json:"money"`
	Introduction string `json:"introduction"`
}
