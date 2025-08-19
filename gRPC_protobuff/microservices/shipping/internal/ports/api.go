package ports

import "github.com/ruandg/microservices/shipping/internal/application/core/domain"

type APIPort interface {
	CalculateDeliveryDays(items []domain.ShippingItem) int32
}
