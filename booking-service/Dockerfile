FROM golang:1.24.2-alpine

WORKDIR /app

# Copy both app and booking-service
COPY app/ ./app/
COPY booking-service/ ./booking-service/

WORKDIR /app/booking-service

RUN go mod tidy
RUN go build -o booking-service .

EXPOSE 8082

CMD ["./booking-service"]