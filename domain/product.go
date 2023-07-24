package domain

type Product struct {
	ID          int
	Name        string
	Description string
	Price       float64
	Amount      int
}

func CreateProduct(
	name, description string,
	price float64,
	amount int,
) Product {
	return Product{
		Name:        name,
		Description: description,
		Price:       price,
		Amount:      amount,
	}
}
