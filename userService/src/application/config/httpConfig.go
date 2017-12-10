package config

import (
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func loadHTTPConfig(application *Application, config Config) (err error) {
	application.http.hostname = config.HTTP.Hostname
	application.http.port = config.HTTP.Port
	return
}

// ListenAndServe starts http server with supplied handler
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
