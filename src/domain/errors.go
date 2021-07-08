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
		newDomainErr(400001, "잘못된 요청입니다."),
	)
)

// 401
var (
	ErrDuplicatedUser = echo.NewHTTPError(
		http.StatusUnauthorized,
		newDomainErr(401001, "이미 존재하는 회원입니다."),
	)
)

// 500
var (
	ErrInternalServerError = echo.NewHTTPError(
		http.StatusInternalServerError,
		newDomainErr(500001, "알 수 없는 에러가 발생하였습니다."),
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

func newDomainErr(code int, message string) *domainErr {
	return &domainErr{
		code,
		message,
	}
}
