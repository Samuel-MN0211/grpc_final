FROM golang:1.24.5-alpine
WORKDIR /app/microservices/shipping
COPY ./microservices/shipping /app/microservices/shipping
COPY ./microservices-proto /app/microservices/shipping/microservices-proto
COPY ./microservices-proto /app/microservices-proto
RUN go mod tidy && go build -o shipping-service ./cmd/main.go
EXPOSE 50053
CMD ["/app/microservices/shipping/shipping-service"]
