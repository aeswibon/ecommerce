package domain

import "errors"

// Address represents an address
type Address struct {
	Name    string `json:"name"`
	Street  string `json:"street"`
	City    string `json:"city"`
	Country string `json:"country"`
	Zip     string `json:"zip"`
}

// NewAddress creates a new address
func NewAddress(name, street, city, country, zip string) (Address, error) {
	addr := Address{
		Name:    name,
		Street:  street,
		City:    city,
		Country: country,
		Zip:     zip,
	}
	if err := addr.Validate(); err != nil {
		return Address{}, err
	}
	return addr, nil
}

// Validate validates the address
func (a Address) Validate() error {
	if len(a.Name) == 0 {
		return errors.New("Empty name")
	}
	if len(a.Street) == 0 {
		return errors.New("Empty street")
	}
	if len(a.City) == 0 {
		return errors.New("Empty city")
	}
	if len(a.Country) == 0 {
		return errors.New("Empty country")
	}
	if len(a.Zip) == 0 {
		return errors.New("Empty zip")
	}
	return nil
}

func (a Address) String() string {
	return a.Name + ", " + a.Street + ", " + a.City + ", " + a.Country + ", " + a.Zip
}

func (a Address) getName() string {
	return a.Name
}

func (a Address) getStreet() string {
	return a.Street
}

func (a Address) getCity() string {
	return a.City
}

func (a Address) getCountry() string {
	return a.Country
}

func (a Address) getZip() string {
	return a.Zip
}
