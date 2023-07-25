package models

import (
	"vrachapi/configs"
)

func UserLogins(userLogin UserLogin, userLoginData UserLoginData) (UserLogin, UserLoginData, error) {
	db := configs.GetDB()
	sqlStatement := "SELECT id, name, email FROM users WHERE username=$1 AND password=$2"
	row := db.QueryRow(sqlStatement, userLogin.UserUsername, userLogin.UserPassword)

	err := row.Scan(
		&userLoginData.UserId,
		&userLoginData.UserName,
		&userLoginData.UserEmail,
	)

	if err != nil {
		return userLogin, userLoginData, err
	}
	return userLogin, userLoginData, nil
}

func GetUserByUsername(username string) (UserLoginData, error) {
	db := configs.GetDB()
	sqlStatement := "SELECT id, name, email FROM users WHERE username=$1"
	row := db.QueryRow(sqlStatement, username)

	userLoginData := UserLoginData{}
	err := row.Scan(
		&userLoginData.UserId,
		&userLoginData.UserName,
		&userLoginData.UserEmail,
	)

	if err != nil {
		return userLoginData, err
	}
	return userLoginData, nil
}
