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
	rc := root.NewRootController()
	e.GET("/", rc.Root)
}

func users(e *echo.Echo) {
	uc := user.NewUserController()
	e.POST("/users", uc.Signup)
	e.POST("/users/self-authenticate", uc.SelfAuthenticate)
}
