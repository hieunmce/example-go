package order

import (
	"context"

	"github.com/jinzhu/gorm"

	"github.com/go-kit/kit/endpoint"
	"github.com/minhkhiemm/example-go/domain"
	"github.com/minhkhiemm/example-go/errorer"
	"github.com/minhkhiemm/example-go/service"
)

func GetAllByDate(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(domain.OrderDate)
		res, err := s.OrderService.GetAllByDate(ctx, req)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return make([]domain.Order, 0), nil
			}

			return nil, err
		}

		return res, nil
	}
}

func Create(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(domain.Order)
		res, err := s.OrderService.Create(ctx, &req)
		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func Get(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetRequest)
		res, err := s.OrderService.Get(ctx, req.ID)
		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func Update(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(domain.Order)
		oldOrder, err := s.OrderService.Get(ctx, req.ID)
		if err != nil {
			return nil, err
		}

		if req.OrderTime != nil {
			oldOrder.OrderTime = req.OrderTime
		}
		if req.ReceiveTime != nil {
			if req.ReceiveTime.Before(*oldOrder.OrderTime) {
				return nil, errorer.ErrInvalidReceiveTime
			}
			oldOrder.ReceiveTime = req.ReceiveTime
		}
		if !req.AccountID.IsZero() {
			oldOrder.AccountID = req.AccountID
		}
		if !req.DetailID.IsZero() {
			oldOrder.DetailID = req.DetailID
		}
		if !req.ShopID.IsZero() {
			oldOrder.ShopID = req.ShopID
		}
		if req.Status != nil {
			oldOrder.Status = req.Status
		}

		res, err := s.OrderService.Update(ctx, oldOrder)
		if err != nil {
			return nil, err
		}

		return res, nil
	}
}
