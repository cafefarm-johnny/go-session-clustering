package user

import (
	"Go-Session-Clustering/src/domain/user/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Signup(c echo.Context) error {
	dto := new(model.UserDTO)
	if err := c.Bind(dto); err != nil {
		return err
	}

	if err := c.Validate(dto); err != nil {
		return err
	}

	if err := signup(dto); err != nil {
		return err
	}

	return c.String(http.StatusCreated, "회원 가입 성공")
}
