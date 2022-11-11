package routes

import (
	"github.com/baguseka01/golang_echo_RESTAPI/controllers"
	"github.com/labstack/echo/v4"
)

func Route() *echo.Echo {
	e := echo.New()

	e.POST("/create", controllers.CreateUser)
	e.GET("/detail/:id", controllers.DetailUser)
	e.PUT("/update/:id", controllers.UpdateUser)
	e.DELETE("/delete/:id", controllers.DeleteUser)
	e.GET("/all", controllers.AllUser)

	return e
}
