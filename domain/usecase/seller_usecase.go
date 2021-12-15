package usecase

import (
	"duck/domain/repository"
	"fmt"
)

type SellerUsecase struct {
	repo repository.SellerRepository
}

func NewSellerUsecase(repo repository.SellerRepository) *SellerUsecase {
	return &SellerUsecase{repo}
}

func (u *SellerUsecase) Sell(consumerID, productID int) error {
	product := u.repo.GetProduct(productID)
	seller := u.repo.GetSeller(product.SellerID)

	fmt.Printf("%v", seller)

	return nil
}
