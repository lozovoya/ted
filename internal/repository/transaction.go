package repository

import (
	"context"
	"fmt"
	"ted/internal/model"

	"github.com/jackc/pgx/v4/pgxpool"
)

type transactionRepo struct {
	pool        *pgxpool.Pool
	userRepo    User
	accountRepo Account
}

func NewTransactionRepo(pool *pgxpool.Pool, userRepo User, accountRepo Account) Transaction {
	return &transactionRepo{pool: pool, userRepo: userRepo, accountRepo: accountRepo}
}

func (t *transactionRepo) Refill(ctx context.Context, dest string, amount int) (model.Transaction, error) {
	var transaction model.Transaction
	var balance int
	var owner string
	var isIdentified bool
	ok, err := t.accountRepo.IsExist(ctx, dest)
	if err != nil {
		return transaction, fmt.Errorf("repository.Refill: %w", err)
	}
	if !ok {
		return transaction, fmt.Errorf("repository.Refill: Wrong destination")
	}
	tx, err := t.pool.Begin(ctx)
	if err != nil {
		return transaction, fmt.Errorf("repository.Refill: %w", err)
	}
	defer tx.Rollback(ctx)
	dbReq := `UPDATE accounts
			SET balance = balance + $1 
			WHERE id=$2
			RETURNING balance, owner`
	err = tx.QueryRow(ctx, dbReq, amount, dest).Scan(&balance, &owner)
	if err != nil {
		return transaction, fmt.Errorf("repository.Refill: %w", err)
	}
	dbReq = `SELECT is_identified FROM users WHERE id=$1`
	err = tx.QueryRow(ctx, dbReq, owner).Scan(&isIdentified)
	if err != nil {
		return transaction, fmt.Errorf("repository.Refill: %w", err)
	}
	if balance > 100_000 {
		tx.Rollback(ctx)
		return transaction, fmt.Errorf("repository.Refill: Wrong balance")
	}
	if (balance > 10_000) && !isIdentified {
		tx.Rollback(ctx)
		return transaction, fmt.Errorf("repository.Refill: Wrong balance")
	}
	dbReq = `INSERT INTO transactions (type, dest, amount) 
			VALUES ((SELECT id FROM transactions_types WHERE type=$1), $2, $3) 
			RETURNING id, dest, amount`
	err = tx.QueryRow(ctx,
		dbReq, "REF",
		dest,
		amount).Scan(&transaction.ID,
		&transaction.Dest,
		&transaction.Amount)
	if err != nil {
		tx.Rollback(ctx)
		return transaction, fmt.Errorf("repository.Refill: %w", err)
	}
	tx.Commit(ctx)
	transaction.Type = "REF"
	return transaction, nil
}

func (t *transactionRepo) Transfer(ctx context.Context, source string, dest string, amount int) (model.Transaction, error) {
	var transaction model.Transaction
	var balance int
	var owner string
	var isIdentified bool
	ok, err := t.accountRepo.IsExist(ctx, source)
	if err != nil {
		return transaction, fmt.Errorf("repository.Transfer: %w", err)
	}
	if !ok {
		return transaction, fmt.Errorf("repository.Transfer: Wrong source")
	}
	ok, err = t.accountRepo.IsExist(ctx, dest)
	if err != nil {
		return transaction, fmt.Errorf("repository.Transfer: %w", err)
	}
	if !ok {
		return transaction, fmt.Errorf("repository.Transfer: Wrong destination")
	}
	tx, err := t.pool.Begin(ctx)
	if err != nil {
		return transaction, fmt.Errorf("repository.Transfer: %w", err)
	}
	defer tx.Rollback(ctx)
	dbReq := `UPDATE accounts
			SET balance = balance - $1 
			WHERE id=$2
			RETURNING balance`
	err = tx.QueryRow(ctx, dbReq, amount, source).Scan(&balance)
	if err != nil {
		return transaction, fmt.Errorf("repository.Transfer: %w", err)
	}
	if balance < 0 {
		tx.Rollback(ctx)
		return transaction, fmt.Errorf("repository.Transfer: Low source balance")
	}
	dbReq = `UPDATE accounts
			SET balance = balance + $1 
			WHERE id=$2
			RETURNING balance, owner`
	err = tx.QueryRow(ctx, dbReq, amount, dest).Scan(&balance, &owner)
	if err != nil {
		return transaction, fmt.Errorf("repository.Transfer: %w", err)
	}
	dbReq = `SELECT is_identified FROM users WHERE id=$1`
	err = tx.QueryRow(ctx, dbReq, owner).Scan(&isIdentified)
	if err != nil {
		return transaction, fmt.Errorf("repository.Transfer: %w", err)
	}
	if balance > 100_000 {
		tx.Rollback(ctx)
		return transaction, fmt.Errorf("repository.Transfer: Wrong balance")
	}
	if (balance > 10_000) && !isIdentified {
		tx.Rollback(ctx)
		return transaction, fmt.Errorf("repository.Transfer: Wrong balance")
	}
	dbReq = `INSERT INTO transactions (type, source, dest, amount) 
			VALUES ((SELECT id FROM transactions_types WHERE type=$1), $2, $3, $4) 
			RETURNING id, source, dest, amount`
	err = tx.QueryRow(ctx,
		dbReq, "TRANS",
		source,
		dest,
		amount).Scan(&transaction.ID,
		&transaction.Source,
		&transaction.Dest,
		&transaction.Amount)
	if err != nil {
		tx.Rollback(ctx)
		return transaction, fmt.Errorf("repository.Transfer: %w", err)
	}
	tx.Commit(ctx)
	transaction.Type = "TRANS"
	return transaction, nil
}

func (t *transactionRepo) GetCurrentMonthTransactionsByAccountID(ctx context.Context, accountID string) ([]model.Transaction, error) {
	var transactions = make([]model.Transaction, 0)
	dbReq := `SELECT  id, source, dest, amount, created 
				FROM transactions 
				WHERE (source=$1 OR dest=$1) AND (date_part('month', created)=date_part('month', now()))`
	rows, err := t.pool.Query(ctx, dbReq, accountID)
	if err != nil {
		return transactions, fmt.Errorf("repository.GetCurrentMonthTransactionsByAccountID: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var transaction model.Transaction
		err = rows.Scan(&transaction.ID,
			&transaction.Source,
			&transaction.Dest,
			&transaction.Amount,
			&transaction.Time)
		if err != nil {
			return transactions, fmt.Errorf("repository.GetCurrentMonthTransactionsByAccountID: %w", err)
		}
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}
