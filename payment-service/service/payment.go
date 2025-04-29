package service

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/rodolfodiazr/booking-system/app/kafka"
	"github.com/rodolfodiazr/booking-system/app/models"
)

// ProcessPayments listens to booking-created events and processes payments
func ProcessPayments() {
	reader := kafka.NewKafkaReader("booking-created", "payment-service-group")

	for {
		message, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Println("Error reading booking created event:", err)
			continue
		}

		var booking models.BookingCreated
		if err := json.Unmarshal(message.Value, &booking); err != nil {
			log.Println("Error unmarshalling booking event:", err)
			continue
		}

		log.Printf("Processing payment for booking ID: %s\n", booking.BookingID)

		// Simulate payment processing
		payment := models.PaymentCompleted{
			PaymentID:  generatePaymentID(),
			BookingID:  booking.BookingID,
			AmountPaid: booking.TotalAmount,
			PaidAt:     time.Now().Format(time.RFC3339),
		}

		// Produce payment completed event
		if err := kafka.ProduceEvent("payment-completed", payment); err != nil {
			log.Println("Error producing payment completed event:", err)
			continue
		}

		log.Printf("Payment processed for booking ID: %s\n", booking.BookingID)
	}
}

// generatePaymentID generates a unique payment ID
func generatePaymentID() string {
	return "PAY56789"
}
