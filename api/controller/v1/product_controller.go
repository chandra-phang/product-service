package v1

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"product-service/api/controller"
	v1request "product-service/dto/request/v1"
	v1response "product-service/dto/response/v1"
	"product-service/infrastructure/log"
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

	dto := new(v1response.ListProductDTO).ConvertFromProductsEntity(products)

	return controller.WriteSuccess(ctx, http.StatusOK, dto)
}

func (c *productController) CreateProduct(ctx echo.Context) error {
	reqBody, _ := ioutil.ReadAll(ctx.Request().Body)
	dto := v1request.CreateProductDTO{}

	if err := json.Unmarshal(reqBody, &dto); err != nil {
		log.Errorf(ctx, err, "[ProductController][CreateProduct] Failed to unmarshal request body %v into dto", reqBody)
		return controller.WriteError(ctx, http.StatusBadRequest, err)
	}

	err := dto.Validate(ctx)
	if err != nil {
		log.Errorf(ctx, err, "[ProductController][CreateProduct] Validation failed for request dto %v ", dto)
		return controller.WriteError(ctx, http.StatusBadRequest, err)
	}

	err = c.svc.CreateProduct(ctx, dto)
	if err != nil {
		log.Errorf(ctx, err, "[ProductController][CreateProduct] Failed to create product for request dto %v ", dto)
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

	dto := new(v1response.GetProductDTO).ConvertFromProductEntity(product)

	return controller.WriteSuccess(ctx, http.StatusOK, dto)
}

func (c *productController) UpdateProduct(ctx echo.Context) error {
	productID := ctx.Param("id")

	reqBody, _ := ioutil.ReadAll(ctx.Request().Body)
	dto := v1request.UpdateProductDTO{}

	if err := json.Unmarshal(reqBody, &dto); err != nil {
		log.Errorf(ctx, err, "[ProductController][UpdateProduct] Failed to unmarshal request body %v into dto", reqBody)
		return controller.WriteError(ctx, http.StatusBadRequest, err)
	}

	if err := json.Unmarshal(reqBody, &dto); err != nil {
		log.Errorf(ctx, err, "[ProductController][UpdateProduct] Failed to unmarshal request body %v into dto", reqBody)
		return controller.WriteError(ctx, http.StatusBadRequest, err)
	}

	err := dto.Validate(ctx)
	if err != nil {
		log.Errorf(ctx, err, "[ProductController][UpdateProduct] Validation failed for request dto %v ", dto)
		return controller.WriteError(ctx, http.StatusBadRequest, err)
	}

	err = c.svc.UpdateProduct(ctx, productID, dto)
	if err != nil {
		log.Errorf(ctx, err, "[ProductController][UpdateProduct] Failed to create product for request dto %v ", dto)
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
	productID := ctx.Param("id")

	err := c.svc.IncreaseBookedQuota(ctx, productID)
	if err != nil {
		return controller.WriteError(ctx, http.StatusInternalServerError, err)
	}

	return controller.WriteSuccess(ctx, http.StatusOK, nil)
}

func (c *productController) DecreaseBookedQuota(ctx echo.Context) error {
	productID := ctx.Param("id")

	err := c.svc.DecreaseBookedQuota(ctx, productID)
	if err != nil {
		return controller.WriteError(ctx, http.StatusInternalServerError, err)
	}

	return controller.WriteSuccess(ctx, http.StatusOK, nil)
}
