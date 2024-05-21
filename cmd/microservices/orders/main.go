package orders

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aeswibon/ecommerce/pkg/common/cmd"
	orders_infra_payments "github.com/aeswibon/ecommerce/pkg/orders/infrastructure/payments"
	"github.com/go-chi/chi"
)

func createOrders() (router *chi.Mux, closeFn func()) {
	cmd.WaitForService(os.Getenv("SHOP_RABBITMQ_ADDR"))
	ordersToPayQueue, err := orders_infra_payments.NewAMQPSvc(
		fmt.Sprintf("amqp://%s/", os.Getenv("SHOP_RABBITMQ_ADDR")),
		os.Getenv("SHOP_RABBITMQ_ORDERS_TO_PAY_QUEUE"),
	)
	if err != nil {
		log.Fatalf("Error creating AMQP service: %s", err)
		panic(err)
	}
	r := cmd.CreateRouter()
	// create routes
	return r, func() {
		if err := ordersToPayQueue.Close(); err != nil {
			log.Fatalf("Error closing AMQP service: %s", err)
		}
	}
}

func main() {
	log.Println("Starting Orders Microservice")
	ctx := cmd.CreateContext()
	r, closeFn := createOrders()
	defer closeFn()
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", os.Getenv("SHOP_ORDERS_PORT")),
		Handler: r,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Panicf("Error starting server: %s", err)
		}
	}()
	<-ctx.Done()
	log.Println("Shutting down Orders Microservice")
	if err := server.Close(); err != nil {
		log.Panicf("Error shutting down server: %s", err)
	}
}
