package main

import (
	"log"

	"github.com/rodolfodiazr/booking-system/app/kafka"
	"github.com/rodolfodiazr/booking-system/notification-service/service"
)

func main() {
	log.Println("Notification service started...")

	kafka.CreateTopic("booking-confirmed")

	service.ProcessNotifications()
}
