package models

import "github.com/golang-jwt/jwt/v5"

type User struct {
	UserId        int    `json:"id"`
	UserName      string `json:"name"`
	UserUsername  string `json:"username"`
	UserEmail     string `json:"email"`
	UserCreatedAt string `json:"created_at"`
	UserUpdatedAt string `json:"updated_at"`
}

type Users struct {
	Users []User `json:"users"`
}

type UserLogin struct {
	UserUsername string `json:"username"`
	UserPassword string `json:"password"`
}

type UserLoginData struct {
	UserId    int    `json:"id"`
	UserName  string `json:"name"`
	UserEmail string `json:"email"`
}

type JwtCustomClaims struct {
	UserUsername string `json:"username"`
	UserName     string `json:"name"`
	UserEmail    string `json:"email"`
	jwt.RegisteredClaims
}
