module github.com/rodolfodiazr/booking-system/payment-service

go 1.24.2

require github.com/rodolfodiazr/booking-system/app v0.0.0

require (
	github.com/klauspost/compress v1.15.9 // indirect
	github.com/pierrec/lz4/v4 v4.1.15 // indirect
	github.com/segmentio/kafka-go v0.4.47 // indirect
)

replace github.com/rodolfodiazr/booking-system/app => ../app
