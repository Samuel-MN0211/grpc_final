FROM golang:1.24.5-alpine
WORKDIR /app/microservices/order
COPY ./microservices/order /app/microservices/order
COPY ./microservices-proto /app/microservices/order/microservices-proto
COPY ./microservices-proto /app/microservices-proto
RUN go mod tidy && go build -o order-service ./cmd/main.go
EXPOSE 50051
CMD ["/app/microservices/order/order-service"]
