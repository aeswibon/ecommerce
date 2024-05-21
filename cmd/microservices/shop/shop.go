package shop

import (
	"log"
	"net/http"
	"os"

	"github.com/aeswibon/ecommerce/pkg/common/cmd"
	"github.com/go-chi/chi"
)

func createShop() *chi.Mux {
	// create repo and service instances
	r := cmd.CreateRouter()
	// add routes
	return r
}

func main() {
	log.Println("Starting shop microservice")
	ctx := cmd.CreateContext()

	r := createShop()
	server := &http.Server{Addr: os.Getenv("SHOP_SHOP_SERVICE_BIND_ADDR"), Handler: r}
	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Panicf("Error starting server: %s", err)
		}
	}()
	<-ctx.Done()
	log.Println("Closing shop microservice")
	if err := server.Close(); err != nil {
		log.Panicf("Error closing server: %s", err)
	}
}
