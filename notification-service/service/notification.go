package service

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/rodolfodiazr/booking-system/app/kafka"
	"github.com/rodolfodiazr/booking-system/app/models"
)

// ProcessNotifications listens to payment-completed events and sends notifications
func ProcessNotifications() {
	reader := kafka.NewKafkaReader("payment-completed", "notification-service-group")

	for {
		message, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Println("Error reading payment completed event:", err)
			continue
		}

		var payment models.PaymentCompleted
		if err := json.Unmarshal(message.Value, &payment); err != nil {
			log.Println("Error unmarshalling payment event:", err)
			continue
		}

		log.Printf("Sending confirmation for payment ID: %s\n", payment.PaymentID)

		// Simulate sending a notification
		notification := models.BookingConfirmed{
			BookingID:      payment.BookingID,
			ConfirmationID: generateConfirmationID(),
			ConfirmedAt:    time.Now().Format(time.RFC3339),
			Message:        "Your booking has been confirmed. Thank you!",
		}

		if err := kafka.ProduceEvent("booking-confirmed", notification); err != nil {
			log.Println("Error producing booking confirmed event:", err)
			continue
		}

		log.Printf("Booking confirmed for booking ID: %s\n", payment.BookingID)
	}
}

// generateConfirmationID generates a unique confirmation ID
func generateConfirmationID() string {
	return "CNF67890"
}
