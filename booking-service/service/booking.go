package service

import (
	"log"

	"github.com/rodolfodiazr/booking-system/app/kafka"
	"github.com/rodolfodiazr/booking-system/app/models"
)

// CreateBooking handles booking creation and produces a BookingCreated event.
func CreateBooking(userID, hotelID string, nights int, totalAmount float64) error {
	bookingID := generateBookingID()

	// Create the booking event
	bookingEvent := models.BookingCreated{
		BookingID:   bookingID,
		UserID:      userID,
		HotelID:     hotelID,
		Nights:      nights,
		TotalAmount: totalAmount,
	}

	// Produce the event to Kafka
	if err := kafka.ProduceEvent("booking-created", bookingEvent); err != nil {
		log.Println("Error producing booking created event:", err)
		return err
	}

	log.Println("Booking created successfully and event produced")
	return nil
}

// generateBookingID is a placeholder for generating unique booking IDs.
func generateBookingID() string {
	return "BKG12345"
}
