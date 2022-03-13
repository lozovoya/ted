package v1

import (
	"ted/internal/repository"

	log "github.com/sirupsen/logrus"
	"github.com/unrolled/render"
)

type Transaction struct {
	transactionRepo repository.Transaction
	lg              *log.Entry
	renderer        render.Render
}

func NewTransactionRepo(transactionRepo repository.Transaction, lg *log.Entry, renderer render.Render) *Transaction {
	return &Transaction{
		transactionRepo: transactionRepo,
		lg:              lg,
		renderer:        renderer,
	}
}
