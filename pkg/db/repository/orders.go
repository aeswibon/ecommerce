package repository

import "github.com/aeswibon/ecommerce/pkg/db/domain"

// OrderRepository interface for order repository
type OrderRepository interface {
	Save(order *domain.Order) error
	FetchByID(id string) (*domain.Order, error)
}
