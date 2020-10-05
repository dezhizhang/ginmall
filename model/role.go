package model

import (
	_ "github.com/jinzhu/gorm"
)

type Role struct {
	Id          int    `json:"id" form:"id"`
	Title       string `json:"title" form:"title"`
	Status      int    `json:"status" form:"status"`
	AddTime     int    `json:"add_time" form:"add_time"`
	Description string `json:"description" form:"description"`
}
