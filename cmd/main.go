package main

import (
	"crypto/tls"
	"net"
	"net/http"
	"os"
	"ted/internal/api/httpserver"
	"time"

	"github.com/caarlos0/env"
	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"
)

type Params = struct {
	//API
	Port    string `env:"TTS_PORT" envDefault:"9999"`
	Host    string `env:"TTS_HOST" envDefault:"0.0.0.0"`
	TLSCert string `env:"TLS_CERT" envDefault:"/etc/tls/tls.crt"`
	TLSKey  string `env:"TLS_KEY" envDefault:"/etc/tls/tls.key"`
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
	//renderer := render.New(render.Options{
	//	DisableHTTPErrorRendering: true,
	//})

	router := httpserver.NewRouter(chi.NewRouter(), lg)
	server := http.Server{
		Addr:        net.JoinHostPort(config.Host, config.Port),
		Handler:     &router,
		IdleTimeout: time.Second * 30,
		TLSConfig:   &tls.Config{MinVersion: tls.VersionTLS13},
	}
	return server.ListenAndServe()
	//return server.ListenAndServeTLS(config.TLSCert, config.TLSKey)
}
