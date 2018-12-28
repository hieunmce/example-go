package account

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/minhkhiemm/example-go/domain"
	"github.com/minhkhiemm/example-go/service"
)

func Create(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(domain.Account)
		res, err := s.AccountService.Create(ctx, &req)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
}
