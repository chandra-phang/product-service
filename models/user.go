package models

import (
	"time"

	"github.com/labstack/echo/v4"
)

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type IUserRepository interface {
	CreateUser(ctx echo.Context, user User) error
}
