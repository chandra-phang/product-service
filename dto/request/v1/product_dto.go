package request

import (
	"errors"
	"product-service/apperrors"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CreateProductDTO struct {
	Name       string `json:"name" validate:"required"`
	DailyQuota int    `json:"dailyQuota" validate:"required"`
}

type UpdateProductDTO struct {
	Name       string `json:"name" validate:"required"`
	DailyQuota int    `json:"dailyQuota" validate:"required"`
}

type IncreaseBookedQuotaDTO struct {
	Products []ProductDTO `json:"products" validate:"required,dive"`
}

type DecreaseBookedQuotaDTO struct {
	Products []ProductDTO `json:"products" validate:"required,dive"`
}

type ProductDTO struct {
	ProductID string `json:"productId" validate:"required"`
	Quantity  int    `json:"quantity" validate:"required"`
}

func (dto CreateProductDTO) Validate(ctx echo.Context) error {
	validate := validator.New()
	if err := validate.Struct(dto); err != nil {
		vErr := apperrors.TryTranslateValidationErrors(err)
		return errors.New(vErr)
	}

	return nil
}

func (dto UpdateProductDTO) Validate(ctx echo.Context) error {
	validate := validator.New()
	if err := validate.Struct(dto); err != nil {
		vErr := apperrors.TryTranslateValidationErrors(err)
		return errors.New(vErr)
	}

	return nil
}

func (dto IncreaseBookedQuotaDTO) Validate(ctx echo.Context) error {
	validate := validator.New()
	if err := validate.Struct(dto); err != nil {
		vErr := apperrors.TryTranslateValidationErrors(err)
		return errors.New(vErr)
	}

	return nil
}

func (dto DecreaseBookedQuotaDTO) Validate(ctx echo.Context) error {
	validate := validator.New()
	if err := validate.Struct(dto); err != nil {
		vErr := apperrors.TryTranslateValidationErrors(err)
		return errors.New(vErr)
	}

	return nil
}
