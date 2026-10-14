package products

import "context"

type Service interface {
	ListProducts(ctx context.Context) error
}

type svc struct {
	// repository Repository
}

func NewService() Service {
	return &svc{
		// repository: NewRepository(),
	}
}

func (s *svc) ListProducts(ctx context.Context) error {
	return nil
}
