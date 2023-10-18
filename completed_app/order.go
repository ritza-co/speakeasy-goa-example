package club

import (
	order "app/gen/order"
	"context"
	"log"
)

// order service example implementation.
// The example methods log the requests and return zero values.
type ordersrvc struct {
	logger *log.Logger
}

// NewOrder returns the order service implementation.
func NewOrder(logger *log.Logger) order.Service {
	return &ordersrvc{logger}
}

// Order a cup of tea.
func (s *ordersrvc) Tea(ctx context.Context, p *order.TeaPayload) (res string, err error) {
	s.logger.Print("order.tea")
	return "", nil
}
