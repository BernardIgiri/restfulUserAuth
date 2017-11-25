package config

import (
	"github.com/justinas/nosurf"
)

func loadSecurityConfig(application *Application, config Config) (err error) {
	application.Middleware = application.Middleware.Append(nosurf.NewPure)
	return
}
