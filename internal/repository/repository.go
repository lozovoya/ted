package repository

import (
	"context"
	"ted/internal/model"
)

type User interface {
	AddUser(ctx context.Context, user model.User) (model.User, error)
	GetUserByID(ctx context.Context, userID string) (model.User, error)
	EditUserByID(ctx context.Context, user model.User) (model.User, error)
	DeleteUserByID(ctx context.Context, user model.User) error
	GetUserKeyByID(ctx context.Context, userID string) (string, error)
}

type Account interface {
	AddAccount(ctx context.Context, account model.Account) (model.Account, error)
	EditAccount(ctx context.Context, account model.Account) (model.Account, error)
	IsExist(ctx context.Context, accountID string) (bool, error)
	GetBalanceByID(ctx context.Context, accountID string) (int, error)
}

type Transaction interface {
	Refill(ctx context.Context, dest string, amount int) (model.Transaction, error)
	Transfer(ctx context.Context, source string, dest string, amount int) (model.Transaction, error)
	GetCurrentMonthTransactionsByAccountID(ctx context.Context, accountID string) ([]model.Transaction, error)
}
