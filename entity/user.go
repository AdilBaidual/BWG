package entity

type User struct {
	Id       int    `json:"id" db:"id"`
	Username string `json:"username" db:"username" binding:"required"`
	Surname  string `json:"surname" db:"surname" binding:"required"`
}
