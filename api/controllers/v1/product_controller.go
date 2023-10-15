package v1

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	controller "product-service/api/controllers"
	v1req "product-service/dto/request/v1"
	v1resp "product-service/dto/response/v1"
	"product-service/services"

	"github.com/labstack/echo/v4"
)

type productController struct {
	svc services.IProductService
}

// creates a new instance of this controller with reference to ProductService
func InitProductController() *productController {
	//  initializes its "svc" field with a service instance returned by "application.GetProductService()".
	return &productController{
		svc: services.GetProductService(),
	}
}

func (c *productController) ListProducts(ctx echo.Context) error {
	products, err := c.svc.ListProducts(ctx)
	if err != nil {
		return controller.WriteError(ctx, http.StatusInternalServerError, err)
	}

	resp := new(v1resp.ListProductDTO).ConvertFromProductsEntity(products)
	return controller.WriteSuccess(ctx, http.StatusOK, resp)
}

func (c *productController) CreateProduct(ctx echo.Context) error {
	reqBody, err := ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		return controller.WriteError(ctx, http.StatusBadRequest, err)
	}

	dto := v1req.CreateProductDTO{}
	if err := json.Unmarshal(reqBody, &dto); err != nil {
		return controller.WriteError(ctx, http.StatusBadRequest, err)
	}

	err = dto.Validate(ctx)
	if err != nil {
		return controller.WriteError(ctx, http.StatusBadRequest, err)
	}

	err = c.svc.CreateProduct(ctx, dto)
	if err != nil {
		return controller.WriteError(ctx, http.StatusInternalServerError, err)
	}

	return controller.WriteSuccess(ctx, http.StatusCreated, nil)
}

func (c *productController) GetProduct(ctx echo.Context) error {
	productID := ctx.Param("id")
	product, err := c.svc.GetProduct(ctx, productID)
	if err != nil {
		return controller.WriteError(ctx, http.StatusInternalServerError, err)
	}

	resp := new(v1resp.GetProductDTO).ConvertFromProductEntity(product)
	return controller.WriteSuccess(ctx, http.StatusOK, resp)
}

func (c *productController) UpdateProduct(ctx echo.Context) error {
	productID := ctx.Param("id")
	reqBody, err := ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		return controller.WriteError(ctx, http.StatusBadRequest, err)
	}

	dto := v1req.UpdateProductDTO{}
	if err := json.Unmarshal(reqBody, &dto); err != nil {
		return controller.WriteError(ctx, http.StatusBadRequest, err)
	}

	err = dto.Validate(ctx)
	if err != nil {
		return controller.WriteError(ctx, http.StatusBadRequest, err)
	}

	err = c.svc.UpdateProduct(ctx, productID, dto)
	if err != nil {
		return controller.WriteError(ctx, http.StatusInternalServerError, err)
	}

	return controller.WriteSuccess(ctx, http.StatusOK, nil)
}

func (c *productController) DisableProduct(ctx echo.Context) error {
	productID := ctx.Param("id")
	err := c.svc.DisableProduct(ctx, productID)
	if err != nil {
		return controller.WriteError(ctx, http.StatusInternalServerError, err)
	}

	return controller.WriteSuccess(ctx, http.StatusOK, nil)
}

func (c *productController) EnableProduct(ctx echo.Context) error {
	productID := ctx.Param("id")
	err := c.svc.EnableProduct(ctx, productID)
	if err != nil {
		return controller.WriteError(ctx, http.StatusInternalServerError, err)
	}

	return controller.WriteSuccess(ctx, http.StatusOK, nil)
}

func (c *productController) IncreaseBookedQuota(ctx echo.Context) error {
	reqBody, err := ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		return controller.WriteError(ctx, http.StatusBadRequest, err)
	}

	dto := v1req.IncreaseBookedQuotaDTO{}
	if err := json.Unmarshal(reqBody, &dto); err != nil {
		return controller.WriteError(ctx, http.StatusBadRequest, err)
	}

	err = dto.Validate(ctx)
	if err != nil {
		return controller.WriteError(ctx, http.StatusBadRequest, err)
	}

	err = c.svc.IncreaseBookedQuota(ctx, dto)
	if err != nil {
		return controller.WriteError(ctx, http.StatusInternalServerError, err)
	}

	return controller.WriteSuccess(ctx, http.StatusOK, nil)
}

func (c *productController) DecreaseBookedQuota(ctx echo.Context) error {
	reqBody, err := ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		return controller.WriteError(ctx, http.StatusBadRequest, err)
	}

	dto := v1req.DecreaseBookedQuotaDTO{}
	if err := json.Unmarshal(reqBody, &dto); err != nil {
		return controller.WriteError(ctx, http.StatusBadRequest, err)
	}

	err = dto.Validate(ctx)
	if err != nil {
		return controller.WriteError(ctx, http.StatusBadRequest, err)
	}

	err = c.svc.DecreaseBookedQuota(ctx, dto)
	if err != nil {
		return controller.WriteError(ctx, http.StatusInternalServerError, err)
	}

	return controller.WriteSuccess(ctx, http.StatusOK, nil)
}
