package models

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email"  validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}
