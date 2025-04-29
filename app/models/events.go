package models

// BookingCreated represents the event when a booking is made.
type BookingCreated struct {
	BookingID   string  `json:"booking_id"`
	UserID      string  `json:"user_id"`
	HotelID     string  `json:"hotel_id"`
	Nights      int     `json:"nights"`
	TotalAmount float64 `json:"total_amount"`
}

// PaymentCompleted represents the event when a payment is successfully processed.
type PaymentCompleted struct {
	PaymentID  string  `json:"payment_id"`
	BookingID  string  `json:"booking_id"`
	AmountPaid float64 `json:"amount"`
	Status     string  `json:"status"`
	PaidAt     string  `json:"payment_date"`
}

// BookingConfirmed represents the event when a booking is confirmed after payment.
type BookingConfirmed struct {
	BookingID      string `json:"booking_id"`
	Status         string `json:"status"`
	ConfirmationID string `json:"confirmation_id"`
	ConfirmedAt    string `json:"confirmed_at"`
	Message        string `json:"message"`
}
