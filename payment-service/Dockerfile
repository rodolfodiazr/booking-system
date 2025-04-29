FROM golang:1.24.2-alpine

WORKDIR /app

# Copy both app and payment-service
COPY app/ ./app/
COPY payment-service/ ./payment-service/

WORKDIR /app/payment-service

RUN go mod tidy
RUN go build -o payment-service .

EXPOSE 8082

CMD ["./payment-service"]