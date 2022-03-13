package repository

import (
	"context"
	"ted/internal/model"

	"github.com/jackc/pgx/v4/pgxpool"
)

type transactionRepo struct {
	pool *pgxpool.Pool
}

func NewTransactionRepo(pool *pgxpool.Pool) Transaction {
	return &transactionRepo{pool: pool}
}

func (t transactionRepo) Refill(ctx context.Context, dest string, amount int) (model.Transaction, error) {
	panic("implement me")
}

func (t transactionRepo) Transfer(ctx context.Context, source string, dest string, amount int) (model.Transaction, error) {
	panic("implement me")
}

func (t transactionRepo) GetCurrentMonthTransactionsByAccountID(ctx context.Context, accountID string) (*[]model.Transaction, error) {
	panic("implement me")
}
