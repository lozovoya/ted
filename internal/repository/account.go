package repository

import (
	"context"
	"ted/internal/model"

	"github.com/jackc/pgx/v4/pgxpool"
)

type accountRepo struct {
	pool *pgxpool.Pool
}

func NewAccountRepo(pool *pgxpool.Pool) Account {
	return &accountRepo{pool: pool}
}

func (a accountRepo) AddAccount(ctx context.Context, account model.Account) (model.Account, error) {
	panic("implement me")
}

func (a accountRepo) EditAccount(ctx context.Context, account model.Account) (model.Account, error) {
	panic("implement me")
}

func (a accountRepo) IsExist(ctx context.Context, accountID string) (bool, error) {
	panic("implement me")
}

func (a accountRepo) GetBalanceByID(ctx context.Context, accountID string) (int, error) {
	panic("implement me")
}
