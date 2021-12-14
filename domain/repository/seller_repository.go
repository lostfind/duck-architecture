package repository

import "duck/domain/model"

type SellerRepository interface {
	GetSeller(id int) model.Seller
}
