package repository

import "duck/domain/model"

type SellerRepository interface {
	GetSeller(id int) model.Seller
	GetConsumer(salesID int) model.Consumer
	GetProduct(id int) model.Product
}
