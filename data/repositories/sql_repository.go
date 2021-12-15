package repositories

import (
	"database/sql"
	"duck/domain/model"
)

type SqlRepo struct {
	db *sql.DB
}

func NewSqlRepo() *SqlRepo {
	return &SqlRepo{}
}

func (r *SqlRepo) GetConsumer(id int) model.Consumer {
	return model.Consumer{}
}
func (r *SqlRepo) LikedProducts(id int) []model.Product {
	return []model.Product{}
}
func (r *SqlRepo) Buy(product model.Product) error {
	return nil
}
func (r *SqlRepo) OrderHistory(id int) []model.Product {
	return []model.Product{}
}
func (r *SqlRepo) GetProduct(id int) model.Product {
	return model.Product{}
}

// func (r *SqlRepo) GetSeller(id int) model.Seller {
// 	return model.Seller{}
// }
func (r *SqlRepo) SaleHistory() []model.Product {
	return nil
}
