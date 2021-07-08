package www

import (
	"Go-Session-Clustering/src/domain/root"

	"github.com/labstack/echo/v4"
)

func Register(e *echo.Echo) {
	index(e)
}

func index(e *echo.Echo) {
	e.GET("/", root.Root)
}
