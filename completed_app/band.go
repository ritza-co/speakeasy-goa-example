package club

import (
	band "app/gen/band"
	"context"
	"log"
)

// band service example implementation.
// The example methods log the requests and return zero values.
type bandsrvc struct {
	logger *log.Logger
}

// NewBand returns the band service implementation.
func NewBand(logger *log.Logger) band.Service {
	return &bandsrvc{logger}
}

// Choose your jazz style.
func (s *bandsrvc) Play(ctx context.Context, p *band.PlayPayload) (err error) {
	s.logger.Print("band.play")
	return
}
