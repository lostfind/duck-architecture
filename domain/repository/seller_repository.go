package repository

import "duck/domain/model"

type SellerRepository interface {
	GetSeller(id int) model.Seller
	GetConsumer(id int) model.Consumer
	GetProduct(id int) model.Product
	SaleHistory() []model.Product
}
