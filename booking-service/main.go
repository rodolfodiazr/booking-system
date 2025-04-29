package main

import (
	"log"

	"github.com/rodolfodiazr/booking-system/app/kafka"
	"github.com/rodolfodiazr/booking-system/booking-service/service"
)

func main() {
	log.Println("Booking service started...")

	kafka.CreateTopic("booking-created")

	service.CreateBooking("USR123", "HOTEL456", 3, 450.0)
}
