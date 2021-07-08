package www

import (
	"fmt"
	"os"

	"github.com/gorilla/sessions"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
)

const (
	port = ":8080"
)

func StartUp(e *echo.Echo) {
	e.Use(
		middleware.Logger(),
		middleware.CSRF(),
		middleware.Gzip(),
		middleware.RequestID(),
		middleware.Recover(),
		session.Middleware(sessions.NewCookieStore([]byte("secret"))),
	)

	if err := e.Start(port); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
