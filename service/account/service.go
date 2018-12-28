package account

import (
	"context"

	"github.com/minhkhiemm/example-go/domain"
)

type Service interface {
	Create(ctx context.Context, account *domain.Account) (*domain.Account, error)
}
