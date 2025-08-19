package api

import "github.com/ruandg/microservices/shipping/internal/application/core/domain"

type API struct{}

func NewAPI() *API {
	return &API{}
}

func (a *API) CalculateDeliveryDays(items []domain.ShippingItem) int32 {
	totalUnits := int32(0)
	for _, item := range items {
		totalUnits += item.Quantity
	}
	days := 1 + (totalUnits-1)/5
	return days
}
