package v1

import (
	"ted/internal/repository"

	log "github.com/sirupsen/logrus"
	"github.com/unrolled/render"
)

type User struct {
	userRepo repository.User
	lg       *log.Entry
	renderer *render.Render
}

func NewUserController(userRepo repository.User, lg *log.Entry, renderer *render.Render) *User {
	return &User{
		userRepo: userRepo,
		lg:       lg,
		renderer: renderer,
	}
}
