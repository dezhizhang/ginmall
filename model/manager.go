package model

import (
	_ "github.com/jinzhu/gorm"
)

type Manager struct {
	Id          int    `json:"id" form:"id"`
	Email       string `json:"email" form:"email"`
	Status      int    `json:"status" form:"status"`
	RoleId      int    `json:"role_id" form:"role_id"`
	IsSuper     int    `json:"is_super" form:"is_super"`
	AddTime     int    `json:"add_time" form:"add_time"`
	Mobile      string `json:"mobile" form:"mobile"`
	Username    string `json:"username" form:"username"`
	Password    string `json:"password" form:"password"`
	Description string `json:"description" form:"description"`
}
