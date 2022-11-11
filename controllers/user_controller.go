package controllers

import (
	"net/http"

	"github.com/baguseka01/golang_echo_RESTAPI/config"
	"github.com/baguseka01/golang_echo_RESTAPI/models"
	"github.com/labstack/echo/v4"
)

func CreateUser(e echo.Context) error {
	user := models.User{}
	e.Bind(&user)

	err := config.DB.Save(&user).Error
	if err != nil {
		return e.JSON(http.StatusBadGateway, err.Error())
	}

	return e.JSON(http.StatusOK, echo.Map{"message": "Success created user", "user": user})
}

func DetailUser(e echo.Context) error {
	var user models.User
	id := e.Param("id")

	err := config.DB.Where("id=?", id).First(&user).Error
	if err != nil {
		return e.JSON(http.StatusBadGateway, err.Error())
	}

	return e.JSON(http.StatusOK, echo.Map{"detai user": user})
}

func UpdateUser(e echo.Context) error {
	user := models.User{}
	e.Bind(&user)
	id := e.Param("id")

	err := config.DB.Where("id=?", id).Updates(&user).Error
	if err != nil {
		return e.JSON(http.StatusBadGateway, err.Error())
	}

	return e.JSON(http.StatusOK, echo.Map{"message": "Success update user", "update user": user})
}

func DeleteUser(e echo.Context) error {
	user := models.User{}
	e.Bind(&user)
	id := e.Param("id")

	err := config.DB.Model(&user).Where("id=?", id).Delete(&user).Error
	if err != nil {
		return e.JSON(http.StatusBadGateway, err.Error())
	}

	return e.JSON(http.StatusOK, echo.Map{"message": "Success delete user"})
}

func AllUser(e echo.Context) error {
	var user []models.User

	err := config.DB.Find(&user).Error
	if err != nil {
		return e.JSON(http.StatusBadGateway, err.Error())
	}

	return e.JSON(http.StatusOK, echo.Map{"all user": user})
}


