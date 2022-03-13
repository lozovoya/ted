package v1

import (
	"ted/internal/repository"

	log "github.com/sirupsen/logrus"
	"github.com/unrolled/render"
)

type Account struct {
	accountRepo repository.Account
	lg          *log.Entry
	renderer    *render.Render
}

func NewAccountController(accountRepo repository.Account, lg *log.Entry, renderer *render.Render) *Account {
	return &Account{
		accountRepo: accountRepo,
		lg:          lg,
		renderer:    renderer,
	}
}
