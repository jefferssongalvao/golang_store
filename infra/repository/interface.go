package repository

import (
	"store/domain"
)

type RepositoryInterface interface {
	ListAll()
	Create(product *domain.Product)
}
