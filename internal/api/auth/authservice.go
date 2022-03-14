package auth

import (
	"context"
	"fmt"
	"net/http"
	"ted/internal/repository"
)

type AuthService struct {
	userRepo repository.User
}

func NewAuthService(userRepo repository.User) *AuthService {
	return &AuthService{userRepo: userRepo}
}

func (auth *AuthService) GetUserKey(ctx context.Context, userID string) (string, error) {
	key, err := auth.userRepo.GetUserKeyByID(ctx, userID)
	if err != nil {
		return key, fmt.Errorf("auth.GetUserKey: %w", err)
	}
	return key, nil
}

func (auth *AuthService) CheckDigest(digest string, userKey string, request *http.Request) (bool, error) {

}
