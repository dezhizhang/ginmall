package model

type Topic struct {
	Id int `json:"id"`
	Title string `json:"title"`
}

type TopicList struct {
	Id int `json:"id" form:"id"`
	Title string `json:"title" form:"title"`
	Page int `json:"page" form:"page"`
	PageSize int `json:"pageSize" form:"pageSize"`
}