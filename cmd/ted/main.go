package main

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"
	"os"
	"ted/internal/api/httpserver"
	v1 "ted/internal/api/v1"
	"ted/internal/repository"
	"time"

	"github.com/caarlos0/env"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"
	"github.com/unrolled/render"
)

type Params = struct {
	//API
	Port    string `env:"TTS_PORT" envDefault:"9999"`
	Host    string `env:"TTS_HOST" envDefault:"0.0.0.0"`
	TLSCert string `env:"TLS_CERT" envDefault:"/etc/tls/tls.crt"`
	TLSKey  string `env:"TLS_KEY" envDefault:"/etc/tls/tls.key"`

	//DB
	DSN string `env:"TED_DB" envDefault:"postgres://app:pass@localhost:5432/teddb"`
}

func main() {
	var config Params
	err := env.Parse(&config)
	if err != nil {
		log.Printf("Config load error: %v", err)
		os.Exit(1)
	}
	if err = execute(config); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func execute(config Params) (err error) {
	loggrus := log.New()
	loggrus.SetFormatter(&log.JSONFormatter{})
	loggrus.SetOutput(os.Stdout)
	loggrus.ReportCaller = true
	lg := loggrus.WithFields(log.Fields{
		"app": "ted",
	})
	renderer := render.New(render.Options{
		DisableHTTPErrorRendering: true,
	})

	userCtx := context.Background()
	userPool, err := pgxpool.Connect(userCtx, config.DSN)
	if err != nil {
		lg.Error(err)
		return err
	}
	userRepo := repository.NewUserRepo(userPool)
	userController := v1.NewUserController(userRepo, lg, renderer)

	accountCtx := context.Background()
	accountPool, err := pgxpool.Connect(accountCtx, config.DSN)
	if err != nil {
		lg.Error(err)
		return err
	}
	accountRepo := repository.NewAccountRepo(accountPool)
	accoutController := v1.NewAccountController(accountRepo, lg, renderer)

	transactionCtx := context.Background()
	transactionPool, err := pgxpool.Connect(transactionCtx, config.DSN)
	if err != nil {
		lg.Error(err)
		return err
	}
	transactionRepo := repository.NewTransactionRepo(transactionPool, userRepo, accountRepo)
	transactionController := v1.NewTransactionController(transactionRepo, lg, renderer)

	router := httpserver.NewRouter(chi.NewRouter(), lg, userController, accoutController, transactionController)
	server := http.Server{
		Addr:        net.JoinHostPort(config.Host, config.Port),
		Handler:     &router,
		IdleTimeout: time.Second * 30,
		TLSConfig:   &tls.Config{MinVersion: tls.VersionTLS13},
	}
	return server.ListenAndServe()
	//return server.ListenAndServeTLS(config.TLSCert, config.TLSKey)
}
