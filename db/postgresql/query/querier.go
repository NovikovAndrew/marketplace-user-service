package query

import "context"

type Querier interface {
	RegisterCustomer(ctx context.Context, args RegisterCustomerParams) (Customer, error)
	RegisterProducer(ctx context.Context, args RegisterProducerParams) (Producer, error)
	RegisterCourier(ctx context.Context, args RegisterCourierParams) (Courier, error)
}

var _ Querier = (*Queries)(nil)
