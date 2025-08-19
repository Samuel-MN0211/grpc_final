package main

import (
	"log"

	"github.com/ruandg/microservices/order/config"
	"github.com/ruandg/microservices/order/internal/adapters/db"
	"github.com/ruandg/microservices/order/internal/adapters/grpc"
	payment_adapter "github.com/ruandg/microservices/order/internal/adapters/payment"
	shipping_adapter "github.com/ruandg/microservices/order/internal/adapters/shipping"
	"github.com/ruandg/microservices/order/internal/application/core/api"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}

	paymentAdapter, err := payment_adapter.NewAdapter(config.GetPaymentServiceUrl())
	if err != nil {
		log.Fatalf("Failed to initialize payment stub. Error: %v", err)
	}

	shippingAdapter, err := shipping_adapter.NewAdapter(config.GetShippingServiceUrl())
	if err != nil {
		log.Fatalf("Failed to initialize shipping stub. Error: %v", err)
	}

	application := api.NewApplication(dbAdapter, paymentAdapter).WithShipping(shippingAdapter)
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()
}
