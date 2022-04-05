package controllers

import (
	"net/http"
	"sqlite3-gorm/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetAllUsers(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := make([]models.User, 0)
		err := db.Find(&user).Error
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, user)
	}
}

func CreareUser(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		u := new(models.User)
		if err := c.Bind(u); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		if err := db.Create(u).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, u)
	}
}

func GetUserByID(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		u := new(models.User)
		if err := db.First(u, id).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, u)
	}
}

func UpdateUserByID(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		u := new(models.User)
		if err := db.First(u, id).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		if err := c.Bind(u); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		if err := db.Save(u).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, u)
	}
}

func DeleteUserByID(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		u := new(models.User)
		if err := db.First(u, id).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		if err := db.Delete(u).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "User deleted",
			"id":      id,
		})
	}
}
