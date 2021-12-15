package main

import (
	"duck/data/repositories"
	"duck/domain/usecase"
	"fmt"
)

func main() {
	sqlRepo := repositories.NewSqlRepo()

	sellerUsecase := usecase.NewSellerUsecase(sqlRepo)
	consumerUsecase := usecase.NewConsumerUsecase(sqlRepo)

	fmt.Println(sellerUsecase.Sell(1, 1))
	fmt.Println(consumerUsecase.Buy(1, 1))
}
