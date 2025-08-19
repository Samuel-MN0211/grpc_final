package api

import (
	"github.com/ruandg/microservices-proto/golang/shipping/microservices-proto/shipping"
	"github.com/ruandg/microservices/order/internal/application/core/domain"
	"github.com/ruandg/microservices/order/internal/ports"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Application struct {
	db       ports.DBPort
	payment  ports.PaymentPort
	shipping ports.ShippingPort
}

func NewApplication(db ports.DBPort, payment ports.PaymentPort) *Application {
	return &Application{
		db:       db,
		payment:  payment,
		shipping: nil, // to be set with WithShipping
	}
}

func (a *Application) WithShipping(shipping ports.ShippingPort) *Application {
	a.shipping = shipping
	return a
}

func (a *Application) PlaceOrder(order domain.Order) (domain.Order, error) {
	var totalQuantity int32
	for _, item := range order.OrderItems {
		totalQuantity += item.Quantity
		exists, err := a.db.ExistsInInventory(item.ProductCode)
		if err != nil {
			order.Status = "Canceled"
			a.db.Save(&order)
			return domain.Order{}, status.Errorf(codes.Internal, "Internal error checking inventory for item %s: %v", item.ProductCode, err)
		}
		if !exists {
			order.Status = "Canceled"
			a.db.Save(&order)
			return domain.Order{}, status.Errorf(codes.NotFound, "Item with product code '%s' does not exist in inventory", item.ProductCode)
		}
	}
	if totalQuantity > 50 {
		order.Status = "Canceled"
		a.db.Save(&order)
		return domain.Order{}, status.Errorf(codes.InvalidArgument, "Order cannot be placed: total quantity (%d items) exceeds maximum allowed limit of 50 items", totalQuantity)
	}

	err := a.db.Save(&order)
	if err != nil {
		order.Status = "Canceled"
		a.db.Save(&order)
		return domain.Order{}, status.Errorf(codes.Internal, "Failed to save order: %v", err)
	}

	paymentErr := a.payment.Charge(order)
	if paymentErr != nil {
		order.Status = "Canceled"
		a.db.Save(&order)
		return domain.Order{}, status.Errorf(codes.FailedPrecondition, "Payment failed: %v", paymentErr)
	}

	order.Status = "Paid"
	a.db.Save(&order)

	if a.shipping != nil {
		var shippingItems []shipping.ShippingItem
		for _, item := range order.OrderItems {
			shippingItems = append(shippingItems, shipping.ShippingItem{
				ProductCode: item.ProductCode,
				Quantity:    item.Quantity,
			})
		}
		days, shipErr := a.shipping.CreateShipping(order.ID, shippingItems)
		if shipErr != nil {
			order.Status = "ShippingFailed"
			a.db.Save(&order)
			return domain.Order{}, status.Errorf(codes.Unavailable, "Shipping service error: %v", shipErr)
		}
		order.DeliveryDays = days
	}

	return order, nil
}
