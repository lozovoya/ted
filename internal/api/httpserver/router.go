package httpserver

import (
	"ted/internal/api/auth"
	"ted/internal/api/httpserver/mw"
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
	authService *auth.AuthService,
) chi.Mux {
	mux.Use(middleware.Logger)
	mux.Route("/api/v1", func(router chi.Router) {
		RouterUser(router, userController, authService, lg)
		RouterAccount(router, accountController, authService, lg)
		RouterTransaction(router, transactionController, authService, lg)
	})
	lg.Info("new router is activated")
	return *mux
}

func RouterUser(router chi.Router, userController *v1.User, authService *auth.AuthService, lg *logrus.Entry) {
	router.With(mw.Auth(authService, lg)).Post("/users/add", userController.AddUser)
	router.With(mw.Auth(authService, lg)).Post("/users/get", userController.GetUserByID)
}
func RouterAccount(router chi.Router, accountController *v1.Account, authService *auth.AuthService, lg *logrus.Entry) {
	router.With(mw.Auth(authService, lg)).Post("/accounts/add", accountController.AddAccount)
	router.With(mw.Auth(authService, lg)).Post("/accounts/isexist", accountController.IsExist)
	router.With(mw.Auth(authService, lg)).Post("/accounts/balance", accountController.GetBalanceByID)
}
func RouterTransaction(router chi.Router, transactionController *v1.Transaction, authService *auth.AuthService, lg *logrus.Entry) {
	router.With(mw.Auth(authService, lg)).Post("/transactions/refill", transactionController.Refill)
	router.With(mw.Auth(authService, lg)).Post("/transactions/transfer", transactionController.Transfer)
	router.With(mw.Auth(authService, lg)).Post("/transactions/get", transactionController.Transactions)
}
