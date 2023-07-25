package controllers

import (
	"fmt"
	"net/http"
	"time"
	"vrachapi/cmd/users/models"
	"vrachapi/configs"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func GetUsers(c echo.Context) error {
	db := configs.GetDB()
	sqlStatement := "SELECT id, name, username, email, created_at, updated_at FROM users ORDER BY id"
	rows, err := db.Query(sqlStatement)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusCreated, err.Error())
	}
	defer rows.Close()
	result := models.Users{}
	for rows.Next() {
		user := models.User{}
		err2 := rows.Scan(
			&user.UserId,
			&user.UserName,
			&user.UserUsername,
			&user.UserEmail,
			&user.UserCreatedAt,
			&user.UserUpdatedAt,
		)
		if err2 != nil {
			return c.JSON(http.StatusCreated, err2.Error())
		}
		result.Users = append(result.Users, user)
	}
	return c.JSON(http.StatusCreated, result)
}

func GetUserByUsername(c echo.Context) error {
	username := c.Param("username")

	userLoginData, err := models.GetUserByUsername(username)

	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			struct {
				Success bool   `json:"success"`
				Message string `json:"message"`
			}{
				Success: false,
				Message: err.Error(),
			},
		)
	}

	return c.JSON(http.StatusOK, userLoginData)
}

func UserLogin(c echo.Context) error {
	userLogins := models.UserLogin{}
	c.Bind(&userLogins)

	userLogins, userLoginData, err := models.UserLogins(
		userLogins,
		models.UserLoginData{},
	)

	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			struct {
				Success bool  `json:"success"`
				Error   error `json:"error"`
			}{
				Success: false,
				Error:   echo.ErrUnauthorized,
			},
		)
	}

	// Set custom claims
	claims := &models.JwtCustomClaims{
		userLogins.UserUsername,
		userLoginData.UserName,
		userLoginData.UserEmail,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	message := struct {
		Success bool   `json:"success"`
		Token   string `json:"token"`
	}{
		Success: true,
		Token:   t,
	}

	return c.JSON(http.StatusCreated, message)
}

func GetUsernameByJwt(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*models.JwtCustomClaims)

	return c.JSON(http.StatusOK, claims)
}
