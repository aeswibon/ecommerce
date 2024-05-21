package domain

import "errors"

var (
	// ErrPriceTooLow is returned when the price is less than 0
	ErrPriceTooLow = errors.New("price must be greater than 0")
	// ErrInvalidCurrency is returned when the currency is invalid
	ErrInvalidCurrency = errors.New("invalid currency")
)

// Price represents a price
type Price struct {
	price    float64
	currency string
}

// NewPrice creates a new price
func NewPrice(price float64, currency string) (Price, error) {
	if price <= 0 {
		return Price{}, ErrPriceTooLow
	}
	if currency == "" || len(currency) != 3 {
		return Price{}, ErrInvalidCurrency
	}
	return Price{
		price:    price,
		currency: currency,
	}, nil
}

// GetPrice returns the price
func (p Price) GetPrice() float64 {
	return p.price
}

// GetCurrency returns the currency
func (p Price) GetCurrency() string {
	return p.currency
}
