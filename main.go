package main

import (
	"echo-mysql-api/controllers"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusCreated, "Welcome to MVC API with Golang and MySQL!")
	})
	e.GET("/allemployees", controllers.GetAllEmployees)
	e.GET("/employee", controllers.GetEmployee)

	e.Logger.Fatal(e.Start(":8080"))
}
