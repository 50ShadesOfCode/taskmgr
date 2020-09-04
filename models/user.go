package models

//User :
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//TableName :
func (User) TableName() string {
	return "users"
}
