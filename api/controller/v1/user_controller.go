package v1

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"shop-api/api/controller"
	v1request "shop-api/dto/request/v1"
	"shop-api/infrastructure/log"
	"shop-api/services"

	"github.com/labstack/echo/v4"
)

type userController struct {
	svc services.IUserService
}

// creates a new instance of this controller with reference to UserService
func InitUserController() *userController {
	//  initializes its "svc" field with a service instance returned by "application.GetUserService()".
	return &userController{
		svc: services.GetUserService(),
	}
}

func (c *userController) CreateUser(ctx echo.Context) error {
	reqBody, _ := ioutil.ReadAll(ctx.Request().Body)
	dto := v1request.CreateUserDTO{}

	if err := json.Unmarshal(reqBody, &dto); err != nil {
		log.Errorf(ctx, err, "[UserController][CreateUser] Failed to unmarshal request body %v into dto", reqBody)
		return controller.WriteError(ctx, http.StatusBadRequest, err)
	}

	if err := json.Unmarshal(reqBody, &dto); err != nil {
		log.Errorf(ctx, err, "[UserController][CreateUser] Failed to unmarshal request body %v into dto", reqBody)
		return controller.WriteError(ctx, http.StatusBadRequest, err)
	}

	err := dto.Validate(ctx)
	if err != nil {
		log.Errorf(ctx, err, "[UserController][CreateUser] Validation failed for request dto %v ", dto)
		return controller.WriteError(ctx, http.StatusBadRequest, err)
	}

	err = c.svc.CreateUser(ctx, dto)
	if err != nil {
		log.Errorf(ctx, err, "[UserController][CreateUser] Failed to create user for request dto %v ", dto)
		return controller.WriteError(ctx, http.StatusInternalServerError, err)
	}

	return controller.WriteSuccess(ctx, http.StatusCreated, nil)
}
