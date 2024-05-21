package domain

import (
	"errors"

	"github.com/google/uuid"
)

// Product represents a product
type Product struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       Price     `json:"price"`
}

// NewProduct creates a new product
func NewProduct(name, description string, price Price) (Product, error) {
	product := Product{
		ID:          uuid.New(),
		Name:        name,
		Description: description,
		Price:       price,
	}
	if err := product.Validate(); err != nil {
		return Product{}, err
	}
	return product, nil
}

// Validate validates the product
func (p Product) Validate() error {
	if len(p.Name) == 0 {
		return errors.New("Empty name")
	}
	return nil
}

// GetID returns the product ID
func (p Product) GetID() uuid.UUID {
	return p.ID
}

// GetName returns the product name
func (p Product) GetName() string {
	return p.Name
}

// GetDescription returns the product description
func (p Product) GetDescription() string {
	return p.Description
}

// GetPrice returns the product price
func (p Product) GetPrice() Price {
	return p.Price
}
