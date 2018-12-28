package http

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"

	"github.com/minhkhiemm/example-go/endpoints/order"

	"github.com/minhkhiemm/example-go/domain"
)

func DecodeGetAllByDateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	dateInt, monthInt, yearInt := 0, 0, 0

	date := r.URL.Query().Get("date")
	if date != "" {
		dateParse, err := strconv.Atoi(date)
		if err != nil {
			return nil, err
		}
		dateInt = dateParse
	}

	month := r.URL.Query().Get("month")
	if month != "" {
		monthParse, err := strconv.Atoi(month)
		if err != nil {
			return nil, err
		}
		monthInt = monthParse
	}

	year := r.URL.Query().Get("year")
	if year != "" {
		yearParse, err := strconv.Atoi(year)
		if err != nil {
			return nil, err
		}
		yearInt = yearParse
	}

	return domain.OrderDate{
		Date:  dateInt,
		Month: monthInt,
		Year:  yearInt,
	}, nil
}

func DecodeCreateOrderRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := domain.Order{}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func DecodeGetOrderByIDRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := order.GetRequest{}

	id, err := domain.UUIDFromString(chi.URLParam(r, "orderid"))
	if err != nil {
		return nil, err
	}

	req.ID = id

	return req, nil
}

func DecodeUpdateOrder(_ context.Context, r *http.Request) (interface{}, error) {
	req := domain.Order{}

	id, err := domain.UUIDFromString(chi.URLParam(r, "orderid"))
	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	req.ID = id

	return req, nil
}
