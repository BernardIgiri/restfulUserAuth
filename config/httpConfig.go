package config

import (
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func loadHttpConfig(application *Application, config Config) (err error) {
	application.http.hostname = config.Http.Hostname
	application.http.port = config.Http.Port
	return
}

func (application *Application) ListenAndServe(handler http.Handler) error {
	addr := application.http.hostname + ":" +
		strconv.Itoa(application.http.port)
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		application.Logger.Info().Msg("Server stopped.")
		os.Exit(1)
	}()
	application.Logger.Info().Msgf("Starting server on host %s and port %d...", application.http.hostname, application.http.port)
	return http.ListenAndServe(addr, handler)
}
