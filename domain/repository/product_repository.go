package repository

import "duck/domain/model"

type ProductRepository interface {
	GetProduct(id int) model.Product
	GetProducts(ids []int) []model.Product
}
