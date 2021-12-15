package usecase

import (
	"duck/domain/model"
	"duck/domain/repository"
	"errors"
)

type ConsumerUsecase struct {
	repo repository.ConsumerRepository
}

func NewConsumerUsecase(repo repository.ConsumerRepository) *ConsumerUsecase {
	return &ConsumerUsecase{repo}
}

func (u *ConsumerUsecase) Buy(consumerID, productID int) error {
	product := u.repo.GetProduct(productID)
	if product.IsSoldout() {
		return errors.New("Soldout")
	}

	if err := u.repo.Buy(product); err != nil {
		return err
	}

	seller := u.repo.GetSeller(product.SellerID)
	u.Notification(seller)

	return nil
}

func (u *ConsumerUsecase) Notification(seller model.Seller) {

}
