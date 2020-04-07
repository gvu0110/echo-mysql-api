package controllers

import (
	"echo-mysql-api/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetAllEmployees(c echo.Context) error {
	result := models.GetAllEmployees()
	return c.JSON(http.StatusOK, result)
}

func GetEmployee(c echo.Context) error {
	userId, err := strconv.ParseInt(c.QueryParam("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	result := models.GetEmployee(userId)
	return c.JSON(http.StatusOK, result)
}
