package repository

import "duck/domain/model"

type ConsumerRepository interface {
	GetConsumer(id int) model.Consumer
	LikedProducts(id int) []model.Product
	Buy(product model.Product) error
	OrderHistory(id int) []model.Product
}
