package user

import (
	"Go-Session-Clustering/src/domain/user/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

type userController struct {
	us *UserService
}

func NewUserController() *userController {
	return &userController{
		us: NewUserService(),
	}
}

func (uc *userController) Signup(c echo.Context) error {
	dto := new(model.UserDTO)
	if err := c.Bind(dto); err != nil {
		return err
	}

	if err := c.Validate(dto); err != nil {
		return err
	}

	if err := uc.us.Signup(dto); err != nil {
		return err
	}

	return c.String(http.StatusCreated, "회원 가입 성공")
}

func (uc *userController) SelfAuthenticate(c echo.Context) error {
	dto := new(model.UserDTO)
	if err := c.Bind(dto); err != nil {
		return err
	}

	if err := c.Validate(dto); err != nil {
		return err
	}

	if err := uc.us.SelfAuthenticate(dto); err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}
