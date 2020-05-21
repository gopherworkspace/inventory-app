package models

import (
	"github.com/dgrijalva/jwt-go"
)

type Register struct {
	UserName string `json:"user_name" validate:"required"`
	Password string `json:"password" validate:"passwd"`
	Email    string `json:"email" validate:"required,email"`
	Address  string `json:"address" validate:"required"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

type RegisterResponse struct {
	StatusCode int
	Body       map[string]interface{}
}

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
