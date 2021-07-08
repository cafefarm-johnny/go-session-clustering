package users

import (
	"Go-Session-Clustering/src/domain"
	"Go-Session-Clustering/src/domain/users/model"
	"Go-Session-Clustering/src/hash"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
)

type UserService struct {
	ur *userRepository
}

func NewUserService() *UserService {
	return &UserService{
		ur: NewUserRepository(),
	}
}

func (us *UserService) Signup(dto *model.UserDTO) error {
	hashPwd, err := hash.ToBcrypt(dto.Password)
	if err != nil {
		return domain.ErrInternalServerError
	}
	return us.ur.Save(dto.Username, hashPwd)
}

func (us *UserService) Signin(c echo.Context, dto *model.UserDTO) error {
	u := us.ur.Find(dto.Username)
	if u == nil {
		return domain.ErrNotFoundUser
	}

	if !hash.CompareBcrypt(u.Password, []byte(dto.Password)) {
		return domain.ErrInvalidPassword
	}

	// save session
	return func(c echo.Context, u model.User) error {
		sess, err := session.Get("my-session", c)
		if err != nil {
			return err
		}

		// 세션 존재 여부에 따라 중복 로그인 처리
		if sess != nil {
			return nil
		}

		sess.Options = &sessions.Options{
			Path:     "/",
			Domain:   "localhost",
			MaxAge:   86400 * 7, // 7일
			Secure:   false,
			HttpOnly: true,
			SameSite: http.SameSiteDefaultMode,
		}

		sess.Values["client"] = model.NewSession(u.UUID, u.Username)
		if err = sess.Save(c.Request(), c.Response()); err != nil {
			return err
		}

		return nil
	}(c, *u)
}

func (us *UserService) SelfAuthenticate(dto *model.UserDTO) error {
	u := us.ur.Find(dto.Username)
	if u == nil {
		return domain.ErrNotFoundUser
	}

	if !hash.CompareBcrypt(u.Password, []byte(dto.Password)) {
		return domain.ErrInvalidPassword
	}

	return nil
}
