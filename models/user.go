package models

type User struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	UserID    string `json:"userid"`
	Password  string `json:"password"`
}
type CreateUser struct {
	FirstName string `json:"firstname" binding:"required"`
	LastName  string `json:"lastname" binding:"required"`
	UserID    string `json:"userid" binding:"required"`
	Password  string `json:"password" binding:"required"`
}
type UpdateUser struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	UserID    string `json:"userid"`
	Password  string `json:"password"`
}

type LoginUser struct {
	ID       uint   `json:"id" binding`
	UserID   string `json:"userid" binding:"required"`
	Password string `json:"password" binding:"required"`
}
