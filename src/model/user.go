package model


type User struct {
    ID   int64 `json:"id"`
    Username string `json:"username"`
	Age  int64 `json:"age"`
}
