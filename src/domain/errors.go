package domain

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// 400
var (
	ErrBadRequest = echo.NewHTTPError(
		http.StatusBadRequest,
		NewDomainErr(400001, "잘못된 요청입니다."),
	)
)

// 401
var (
	ErrDuplicatedUser = echo.NewHTTPError(
		http.StatusUnauthorized,
		NewDomainErr(401001, "이미 존재하는 회원입니다."),
	)
)

type domainErr struct {
	Code    int
	Message string
}

// error 인터페이스 덕타이핑
func (d *domainErr) Error() string {
	return fmt.Sprintf("{ code: %d, message: %s }", d.Code, d.Message)
}

func NewDomainErr(code int, message string) *domainErr {
	return &domainErr{
		code,
		message,
	}
}
