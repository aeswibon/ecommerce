package payments

import (
	"fmt"
	"log"
	"os"

	"github.com/aeswibon/ecommerce/pkg/common/cmd"
	"github.com/aeswibon/ecommerce/pkg/payments/interaces/amqp"
)

func createPayments() amqp.PaymentsInterface {
	cmd.WaitForService(os.Getenv("SHOP_RABBITMQ_ADDR"))
	paymentsInterface, err := amqp.NewPaymentsInterface(
		fmt.Sprintf("amqp://%s/", os.Getenv("SHOP_RABBITMQ_ADDR")),
		os.Getenv("SHOP_RABBITMQ_ORDERS_TO_PAY_QUEUE"),
	)
	if err != nil {
		panic(err)
	}
	return paymentsInterface
}

func main() {
	log.Println("Starting payments microservice")
	defer log.Println("Closing payments microservice")
	ctx := cmd.CreateContext()
	paymentsInterface := createPayments()
	if err := paymentsInterface.Run(ctx); err != nil {
		log.Println(err)
	}
}
