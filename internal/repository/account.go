package repository

import (
	"context"
	"errors"
	"fmt"
	"ted/internal/model"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type accountRepo struct {
	pool *pgxpool.Pool
}

func NewAccountRepo(pool *pgxpool.Pool) Account {
	return &accountRepo{pool: pool}
}

func (a *accountRepo) AddAccount(ctx context.Context, account model.Account) (model.Account, error) {
	dbReq := `INSERT INTO accounts (owner, is_active) 
			VALUES ($1, $2) 
			RETURNING id, owner, is_active`
	var addedAccount model.Account
	err := a.pool.QueryRow(ctx,
		dbReq,
		account.Owner,
		account.IsActive).Scan(&addedAccount.ID,
		&addedAccount.Owner,
		&addedAccount.IsActive)
	if err != nil {
		return addedAccount, err
	}
	return addedAccount, nil
}

func (a *accountRepo) EditAccount(ctx context.Context, account model.Account) (model.Account, error) {
	var editedAccount model.Account
	return editedAccount, errors.New("not implemented")
}

func (a *accountRepo) IsExist(ctx context.Context, accountID string) (bool, error) {
	dbReq := `SELECT EXISTS (SELECT * FROM accounts WHERE id=$1)`
	var ok = false
	row := a.pool.QueryRow(ctx, dbReq, accountID)
	if err := row.Scan(&ok); err != nil {
		if err == pgx.ErrNoRows {
			return ok, nil
		}
		return ok, fmt.Errorf("repository.IsExist: %w", err)
	}
	if ok == true {
		return ok, nil
	}
	return false, nil
}

func (a *accountRepo) GetBalanceByID(ctx context.Context, accountID string) (int, error) {
	dbReq := `SELECT balance FROM accounts WHERE id=$1`
	var balance int
	row := a.pool.QueryRow(ctx, dbReq, accountID)
	switch err := row.Scan(&balance); err {
	case pgx.ErrNoRows:
		return balance, errors.New("repository.GetBalanceByID: Wrong ID")
	case nil:
		return balance, nil
	default:
		return balance, errors.New("repository.GetBalanceByID: Internal error")
	}
}
