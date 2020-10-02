package model


type User struct {
    ID   int `json:"id"`
    Username string `json:"username"`
	Age  int64 `json:"age"`
}
