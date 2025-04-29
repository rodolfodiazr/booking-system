package main

import (
	"log"

	"github.com/rodolfodiazr/booking-system/app/kafka"
	"github.com/rodolfodiazr/booking-system/payment-service/service"
)

func main() {
	log.Println("Payment service started...")

	kafka.CreateTopic("payment-completed")

	service.ProcessPayments()
}
