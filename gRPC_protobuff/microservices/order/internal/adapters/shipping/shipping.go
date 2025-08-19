package shipping_adapter

import (
	"context"
	"time"

	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"github.com/ruandg/microservices-proto/golang/shipping/microservices-proto/shipping"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

type Adapter struct {
	client shipping.ShippingClient
}

func NewAdapter(shippingServiceUrl string) (*Adapter, error) {
	var opts []grpc.DialOption
	opts = append(opts,
		grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(
			grpc_retry.WithCodes(codes.Unavailable, codes.ResourceExhausted),
			grpc_retry.WithMax(5),
			grpc_retry.WithBackoff(grpc_retry.BackoffLinear(time.Second)),
		)),
	)
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(shippingServiceUrl, opts...)
	if err != nil {
		return nil, err
	}
	client := shipping.NewShippingClient(conn)
	return &Adapter{client: client}, nil
}

func (a *Adapter) CreateShipping(purchaseID int64, items []shipping.ShippingItem) (int32, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	shippingItems := make([]*shipping.ShippingItem, len(items))
	for i := range items {
		shippingItems[i] = &items[i]
	}
	req := &shipping.CreateShippingRequest{
		PurchaseId: purchaseID,
		Items:      shippingItems,
	}

	resp, err := a.client.Create(ctx, req)
	if err != nil {
		if grpcErr, ok := status.FromError(err); ok && grpcErr.Code() == codes.DeadlineExceeded {
			return 0, err
		}
		return 0, err
	}
	return resp.DeliveryDays, nil
}
