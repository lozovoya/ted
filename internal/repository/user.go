package repository

import (
	"context"
	"fmt"
	"ted/internal/model"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type userRepo struct {
	pool *pgxpool.Pool
}

func NewUserRepo(pool *pgxpool.Pool) User {
	return &userRepo{pool: pool}
}

func (u *userRepo) AddUser(ctx context.Context, user model.User) (model.User, error) {
	dbReq := `INSERT INTO users (name, password, is_identified, is_active) 
			VALUES($1, $2, $3, $4) 
			RETURNING id, name`
	var addedUser model.User
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return addedUser, fmt.Errorf("repository.Adduser: %w", err)
	}
	err = u.pool.QueryRow(ctx,
		dbReq,
		user.Name,
		hash,
		user.IsIdentified,
		user.IsActive).Scan(&addedUser.ID, &addedUser.Name)
	if err != nil {
		return addedUser, fmt.Errorf("repository.Adduser: %w", err)
	}
	return addedUser, nil
}

func (u *userRepo) GetUserByID(ctx context.Context, userID string) (model.User, error) {
	dbReq := `SELECT name, is_identified, is_active 
			FROM users 
			WHERE id=$1`
	var addedUser model.User
	row := u.pool.QueryRow(ctx, dbReq, userID)
	switch err := row.Scan(&addedUser.Name, &addedUser.IsIdentified, &addedUser.IsActive); err {
	case pgx.ErrNoRows:
		return addedUser, nil
	case nil:
		addedUser.ID = userID
		return addedUser, nil
	default:
		return addedUser, nil
	}
}

func (u *userRepo) EditUserByID(ctx context.Context, user model.User) (model.User, error) {
	panic("implement me")
}

func (u *userRepo) DeleteUserByID(ctx context.Context, user model.User) error {
	panic("implement me")
}

func (u *userRepo) GetUserKeyByID(ctx context.Context, userID string) (string, error) {
	panic("implement me")
}
