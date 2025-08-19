package grpc_adapter

import (
	"context"

	"github.com/ruandg/microservices-proto/golang/shipping/microservices-proto/shipping"
	"github.com/ruandg/microservices/shipping/internal/application/core/api"
	"github.com/ruandg/microservices/shipping/internal/application/core/domain"
)

type Server struct {
	shipping.UnimplementedShippingServer
	api *api.API
}

func NewServer(api *api.API) *Server {
	return &Server{api: api}
}

func (s *Server) Create(ctx context.Context, req *shipping.CreateShippingRequest) (*shipping.CreateShippingResponse, error) {
	var items []domain.ShippingItem
	for _, item := range req.Items {
		items = append(items, domain.ShippingItem{
			ProductCode: item.ProductCode,
			Quantity:    item.Quantity,
		})
	}
	days := s.api.CalculateDeliveryDays(items)
	return &shipping.CreateShippingResponse{DeliveryDays: days}, nil
}
