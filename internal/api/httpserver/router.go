package httpserver

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sirupsen/logrus"
)

func NewRouter(
	mux *chi.Mux,
	lg *logrus.Entry,
) chi.Mux {
	mux.Use(middleware.Logger)
	mux.Route("/api/v1", func(router chi.Router) {
		RouterUser(router)
		RouterAccount(router)
		RouterTransaction(router)
	})
	lg.Info("new router is activated")
	return *mux
}

func RouterUser(router chi.Router)        {}
func RouterAccount(router chi.Router)     {}
func RouterTransaction(router chi.Router) {}
