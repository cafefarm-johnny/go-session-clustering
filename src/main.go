package main

import (
	"Go-Session-Clustering/src/www"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	www.Register(e)
	www.StartUp(e)
}
