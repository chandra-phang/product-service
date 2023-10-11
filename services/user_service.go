package services

import (
	v1request "shop-api/dto/request/v1"
	"shop-api/handlers"
	"shop-api/lib"
	"shop-api/models"
	"shop-api/repositories"
	"time"

	"github.com/labstack/echo/v4"
)

type IUserService interface {
	// svc CRUD methods for domain objects
	CreateUser(ctx echo.Context, dto v1request.CreateUserDTO) error
}

type userSvc struct {
	UserRepo models.IUserRepository
}

var userSvcSingleton IUserService

func InitUserService(h handlers.Handler) {
	userSvcSingleton = userSvc{
		UserRepo: repositories.NewJobRepositoryInstance(h.DB),
	}
}

func GetUserService() IUserService {
	return userSvcSingleton
}

func (svc userSvc) CreateUser(ctx echo.Context, dto v1request.CreateUserDTO) error {
	user := models.User{
		ID:        lib.GenerateUUID(),
		Name:      dto.Name,
		Email:     dto.Email,
		Password:  dto.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := svc.UserRepo.CreateUser(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
