package www

import (
	"fmt"
	"os"
	"time"

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
		middleware.CSRF(),
		middleware.Gzip(),
		middleware.LoggerWithConfig(loggerConfig()),
		middleware.RequestID(),
		middleware.Recover(),
		middleware.TimeoutWithConfig(timeoutConfig()),
		session.Middleware(sessionCookieStore()),
	)

	e.Validator = newArgumentValidator()

	if err := e.Start(port); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func loggerConfig() middleware.LoggerConfig {
	return middleware.LoggerConfig{
		Format: "{\n    " +
			"timestamp: \"${time_rfc3339}\" \n    " +
			"request-id: \"${id}\" \n    " +
			"remote-ip: \"${remote_ip}\" \n    " +
			"user-agent: \"${user_agent}\" \n    " +
			"method: \"${method}\" \n    " +
			"path: \"${path}\" \n    " +
			"status: ${status} \n    " +
			"error: \"${error}\" \n" +
			"}",
	}
}

func sessionCookieStore() *sessions.CookieStore {
	return sessions.NewCookieStore([]byte("secret"))
}

func timeoutConfig() middleware.TimeoutConfig {
	return middleware.TimeoutConfig{
		Timeout: 30 * time.Second,
	}
}
