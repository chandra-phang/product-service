package request

import (
	"errors"
	"shop-api/apperrors"
	"shop-api/infrastructure/log"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CreateUserDTO struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (dto CreateUserDTO) Validate(ctx echo.Context) error {
	validate := validator.New()
	if err := validate.Struct(dto); err != nil {
		vErr := apperrors.TryTranslateValidationErrors(err)
		log.Infof(ctx, "[CreateUserDTO] [Validate] Request DTO validation failed %v",
			map[string]interface{}{
				"error":   vErr,
				"request": dto,
			})
		return errors.New(vErr)
	}

	return nil
}
