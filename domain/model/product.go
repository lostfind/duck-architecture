package model

type Product struct {
	ID          int
	Name        string
	ModelNumber string
	SellerID    int
}

func (p *Product) IsSoldout() bool {
	return true
}
