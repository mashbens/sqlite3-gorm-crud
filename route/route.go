package route

import (
	"sqlite3-gorm/config"
	"sqlite3-gorm/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	db := config.InitDB()
	e := echo.New()
	e.GET("/users", controllers.GetAllUsers(db))
	e.POST("/users", controllers.CreareUser(db))
	e.GET("/users/:id", controllers.GetUserByID(db))
	e.PUT("/users/:id", controllers.UpdateUserByID(db))
	e.DELETE("/users/:id", controllers.DeleteUserByID(db))
	return e
}
