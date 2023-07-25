package controllers

import (
	"carilokak/cmd/objects/models"
	"carilokak/configs"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// e.POST("/object/create", objects.CreateObject)
func CreateObject(c echo.Context) error {
	newobject := models.Object{}
	c.Bind(&newobject)
	newObject, err := models.CreateObject(newobject)
	//newObject := new(Objects)
	//newObject, err := models.CreateObject(newobject)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	message := struct {
		Success bool          `json:"success"`
		Data    models.Object `json:"data"`
	}{
		Success: true,
		Data:    newObject,
	}

	return c.JSON(http.StatusCreated, message)
}

func GetObjects(c echo.Context) error {
	db := configs.GetDB()
	sqlStatement := "SELECT object_id, object_name, gender_id FROM objects ORDER BY object_id"
	rows, err := db.Query(sqlStatement)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusCreated, err.Error())
	}
	defer rows.Close()
	result := models.Objects{}
	for rows.Next() {
		object := models.Object{}
		err2 := rows.Scan(&object.ObjectId, &object.ObjectName, &object.ObjectGender)
		if err2 != nil {
			return c.JSON(http.StatusCreated, err2.Error())
		}
		result.Objects = append(result.Objects, object)
	}
	return c.JSON(http.StatusCreated, result)
}

func UpdateObject(c echo.Context) error {
	editObject := models.Object{}
	c.Bind(&editObject)
	editObject, err := models.UpdateObject(editObject)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	message := struct {
		Success bool          `json:"success"`
		Message string        `json:"message"`
		Data    models.Object `json:"data"`
	}{
		Success: true,
		Message: fmt.Sprintf("#%d has been updated successfully.", editObject.ObjectId),
		Data:    editObject,
	}
	return c.JSON(http.StatusCreated, message)
}

func DeleteObject(c echo.Context) error {
	objectId := c.Param("id")
	id := models.DeleteObject(objectId)
	message := struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}{
		Success: true,
		Message: id + " has been deleted successfully.",
	}
	return c.JSON(http.StatusOK, message)
}

func GetData(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "success"})
}
