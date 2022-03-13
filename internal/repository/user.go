package repository

import (
	"context"
	"ted/internal/model"

	"github.com/jackc/pgx/v4/pgxpool"
)

type userRepo struct {
	pool *pgxpool.Pool
}

func NewUserRepo(pool *pgxpool.Pool) User {
	return &userRepo{pool: pool}
}

func (u userRepo) AddUser(ctx context.Context, user model.User) (model.User, error) {
	panic("implement me")
}

func (u userRepo) GetUserByID(ctx context.Context, userID string) (model.User, error) {
	panic("implement me")
}

func (u userRepo) EditUserByID(ctx context.Context, user model.User) (model.User, error) {
	panic("implement me")
}

func (u userRepo) DeleteUserByID(ctx context.Context, user model.User) error {
	panic("implement me")
}
