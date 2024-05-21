package shop

// HTTPClient is a struct that holds the address
type HTTPClient struct {
	address string
}

// NewHTTPClient is a constructor for HTTPClient
func NewHTTPClient(address string) HTTPClient {
	return HTTPClient{address}
}