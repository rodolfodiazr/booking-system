FROM golang:1.24.2-alpine

WORKDIR /app

# Copy both app and notification-service
COPY app/ ./app/
COPY notification-service/ ./notification-service/

WORKDIR /app/notification-service

RUN go mod tidy
RUN go build -o notification-service .

EXPOSE 8082

CMD ["./notification-service"]