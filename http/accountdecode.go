package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/minhkhiemm/example-go/domain"
)

func DecodeCreateAccount(_ context.Context, r *http.Request) (interface{}, error) {
	req := domain.Account{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}
