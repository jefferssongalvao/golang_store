package repository

import (
	"store/domain"
	"store/infra/database"
)

type ProductsRepository struct{}

func (repository *ProductsRepository) ListAll() []domain.Product {
	db := database.Connect()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM products ORDER BY id asc")
	if err != nil {
		panic(err.Error())
	}

	product := domain.Product{}
	products := []domain.Product{}
	for rows.Next() {
		err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Amount)
		if err != nil {
			panic(err.Error())
		}
		products = append(products, product)
	}
	return products
}

func (repository *ProductsRepository) CreateProduct(product *domain.Product) {
	db := database.Connect()
	defer db.Close()

	prepareQuery, err := db.Prepare("INSERT INTO products(name, description, price, amount) VALUES($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	prepareQuery.Exec(product.Name, product.Description, product.Price, product.Amount)
}

func (repository *ProductsRepository) Update(product *domain.Product) {
	db := database.Connect()
	defer db.Close()

	prepareQuery, err := db.Prepare("UPDATE products SET name = $1, description = $2, price = $3, amount = $4 WHERE id = $5")
	if err != nil {
		panic(err.Error())
	}

	prepareQuery.Exec(product.Name, product.Description, product.Price, product.Amount, product.ID)
}

func (repository *ProductsRepository) Get(ID int) domain.Product {
	db := database.Connect()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM products WHERE id = $1", ID)
	if err != nil {
		panic(err.Error())
	}

	product := domain.Product{}
	for rows.Next() {
		err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Amount)
		if err != nil {
			panic(err.Error())
		}
	}

	return product
}

func (repository *ProductsRepository) Delete(ID int) {
	db := database.Connect()
	defer db.Close()

	prepareQuery, err := db.Prepare("DELETE FROM products WHERE id = $1")
	if err != nil {
		panic(err.Error())
	}

	prepareQuery.Exec(ID)
}
