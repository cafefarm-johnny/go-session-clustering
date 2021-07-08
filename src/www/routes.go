package www

import (
	"Go-Session-Clustering/src/domain/root"
	"Go-Session-Clustering/src/domain/user"

	"github.com/labstack/echo/v4"
)

func Register(e *echo.Echo) {
	index(e)
	users(e)
}

func index(e *echo.Echo) {
	e.GET("/", root.Root)
}

func users(e *echo.Echo) {
	e.POST("/users", user.Signup)
}
