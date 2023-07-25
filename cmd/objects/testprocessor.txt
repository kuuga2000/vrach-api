package objects

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// e.POST("/object/create", objects.CreateObject)
func CreateObject(c echo.Context) error {
	return c.JSON(http.StatusOK, "Create an Object")
}

// e.POST("/object/getall", objects.GetAllObjects)
func GetAllObjects(c echo.Context) error {
	return c.JSON(http.StatusOK, "Get All Objects")
}

// e.POST("/object/detail/:id", objects.CreateObject)
func GetObject(c echo.Context) error {
	return c.JSON(http.StatusOK, "Get Object")
}

// e.POST("/object/edit/:id", objects.EditObject)
func EditObject(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, id)
}

// e.POST("/object/delete/:id", objects.CreateObject)
func DeleteObject(c echo.Context) error {
	return c.JSON(http.StatusOK, "Delete")
}

func Message(c echo.Context) error {
	return c.JSON(
		http.StatusOK,
		struct {
			Success bool   `json:"success"`
			Data    string `json:"data"`
		}{
			Success: true,
			Data:    "Data",
		},
	)
}
