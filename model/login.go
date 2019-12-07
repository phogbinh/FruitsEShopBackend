package model

// A Login contains an user login email and password.
type Login struct {
	Mail     string `form:"mail" json:"mail" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
