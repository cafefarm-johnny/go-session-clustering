package root

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func Root(c echo.Context) error {
	sess, err := session.Get("my-session", c)
	if err != nil {
		return err
	}

	sess.Options = &sessions.Options{
		Path:     "/",
		Domain:   "localhost",
		MaxAge:   86400 * 7, // 7일
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteDefaultMode,
	}

	sess.Values["client"] = "johnny"
	if err = sess.Save(c.Request(), c.Response()); err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}
