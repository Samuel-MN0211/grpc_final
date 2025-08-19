package ports

import "github.com/ruandg/microservices-proto/golang/shipping/microservices-proto/shipping"

type ShippingPort interface {
	CreateShipping(purchaseID int64, items []shipping.ShippingItem) (int32, error)
}
