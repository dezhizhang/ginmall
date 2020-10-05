package model

import (
	_ "github.com/jinzhu/gorm"
)

type Manager struct {
	Id          int    `json:"id" form:"id"`
	Email       string `json:"email" form:"email"`
	Status      int    `json:"status" form:"status"`
	RoleId      int    `json:"role_id"`
	IsSuper     int    `json:"is_super"`
	AddTime     int    `json:"add_time"`
	Mobile      string `json:"mobile"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Description string `json:"description"`
}
