package usecase

import (
	"duck/domain/model"
	"duck/domain/repository"
	"errors"
)

type ConsumerUsecase struct {
	consumerRepo repository.ConsumerRepository
	productRepo  repository.ProductRepository
	sellerRepo   repository.SellerRepository
}

func (u *ConsumerUsecase) Buy(consumerID, productID int) error {
	product := u.productRepo.GetProduct(productID)
	if product.IsSoldout() {
		return errors.New("Soldout")
	}

	if err := u.consumerRepo.Buy(product); err != nil {
		return err
	}

	seller := u.sellerRepo.GetSeller(product.SellerID)
	u.Notification(seller)

	return nil
}

func (u *ConsumerUsecase) Notification(seller model.Seller) {

}
