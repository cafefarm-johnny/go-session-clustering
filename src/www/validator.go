package www

import (
	"Go-Session-Clustering/src/domain"

	"github.com/go-playground/validator"
)

type argumentValidator struct {
	validator *validator.Validate
}

func newArgumentValidator() *argumentValidator {
	return &argumentValidator{
		validator: validator.New(),
	}
}

func (av *argumentValidator) Validate(i interface{}) error {
	if err := av.validator.Struct(i); err != nil {
		return domain.ErrBadRequest
	}

	return nil
}
