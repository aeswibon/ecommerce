package repository

import (
	"errors"

	"github.com/aeswibon/ecommerce/pkg/db/domain"
	"github.com/google/uuid"
)

// ErrNotFound is returned when the product is not found
var ErrNotFound = errors.New("product not found")

// ProductRepository represents a repository for a product
type ProductRepository interface {
	Save(*domain.Product) error
	FetchByID(uuid.UUID) (*domain.Product, error)
}
