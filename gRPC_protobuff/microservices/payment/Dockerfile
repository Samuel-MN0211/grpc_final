FROM golang:1.24.5-alpine
WORKDIR /app/microservices/payment
COPY ./microservices/payment /app/microservices/payment
COPY ./microservices-proto /app/microservices/payment/microservices-proto
COPY ./microservices-proto /app/microservices-proto
RUN go mod tidy && go build -o payment-service ./cmd/main.go
EXPOSE 50052
CMD ["/app/microservices/payment/payment-service"]
