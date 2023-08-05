package handler

import (
	"net/http"
	"vrachapi/cmd/customer/model"

	"github.com/labstack/echo/v4"
)

func Register(c echo.Context) error {
	customer := model.Customer{}
	c.Bind(&customer)
	customer, err := model.CreateAccount(customer)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	message := struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}{
		Success: true,
		Message: "Account has been created successfully",
	}

	return c.JSON(http.StatusCreated, message)
}
