package dto

import (
	api "snapp/api/productController"
	"snapp/internal/entity"
)

func product(prs ...api.Product) []entity.Product {
	out := make([]entity.Product, 0)
	for _, val := range prs {
		out = append(out, entity.Product{
			ID:    val.Id,
			Name:  val.Name,
			Point: val.Point,
		})
	}
	return out
}
