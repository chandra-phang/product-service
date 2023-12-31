package v1

import (
	"product-service/lib"
	"product-service/model"
)

type GetProductDTO struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	DailyQuota int    `json:"dailyQuota"`
	Status     string `json:"status"`
	CreatedAt  int64  `json:"createdAt"`
	UpdatedAt  int64  `json:"updatedAt"`
}

type ListProductDTO struct {
	Products []GetProductDTO `json:"products"`
}

func (dto *GetProductDTO) ConvertFromProductEntity(entity *model.Product) *GetProductDTO {
	return &GetProductDTO{
		ID:         entity.ID,
		Name:       entity.Name,
		DailyQuota: entity.DailyQuota,
		Status:     string(entity.Status),
		CreatedAt:  lib.ConvertToEpoch(entity.CreatedAt),
		UpdatedAt:  lib.ConvertToEpoch(entity.UpdatedAt),
	}
}

func (dto *ListProductDTO) ConvertFromProductsEntity(entities []model.Product) *ListProductDTO {
	resp := &ListProductDTO{}
	for _, entity := range entities {
		article := new(GetProductDTO).ConvertFromProductEntity(&entity)
		resp.Products = append(resp.Products, *article)
	}

	return resp
}
