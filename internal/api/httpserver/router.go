package httpserver

import (
	v1 "ted/internal/api/v1"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sirupsen/logrus"
)

func NewRouter(
	mux *chi.Mux,
	lg *logrus.Entry,
	userController *v1.User,
	accountController *v1.Account,
	transactionController *v1.Transaction,
) chi.Mux {
	mux.Use(middleware.Logger)
	mux.Route("/api/v1", func(router chi.Router) {
		RouterUser(router, userController)
		RouterAccount(router, accountController)
		RouterTransaction(router, transactionController)
	})
	lg.Info("new router is activated")
	return *mux
}

func RouterUser(router chi.Router, userController *v1.User) {
	router.Post("/users/add", userController.AddUser)
	router.Post("/users/get", userController.GetUserByID)
}
func RouterAccount(router chi.Router, accountController *v1.Account) {
	router.Post("/accounts/add", accountController.AddAccount)
	router.Post("/accounts/isexist", accountController.IsExist)
	router.Post("/accounts/balance", accountController.GetBalanceByID)
}
func RouterTransaction(router chi.Router, transactionController *v1.Transaction) {
	router.Post("/transactions/refill", transactionController.Refill)
	router.Post("/transactions/transfer", transactionController.Transfer)
	router.Post("/transactions/get", transactionController.Transactions)
}
