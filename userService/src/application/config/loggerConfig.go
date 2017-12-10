package config

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/hlog"
)

func loadLoggerConfig(application *Application, config Config) (err error) {
	switch strings.ToLower(config.Log.Level) {
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case "fatal":
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	case "panic":
		zerolog.SetGlobalLevel(zerolog.PanicLevel)
	case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	}
	errorlogFileHandler, err := os.OpenFile(config.Log.Path,
		os.O_RDWR|os.O_CREATE|os.O_APPEND, 0660)
	if err != nil {
		return err
	}
	application.Logger = zerolog.New(errorlogFileHandler).With().
		Timestamp().
		Str("host", config.HTTP.Hostname).
		Int("port", config.HTTP.Port).
		Logger()
	m := application.Middleware
	// Install the logger handler with default output on the console
	m = m.Append(hlog.NewHandler(application.Logger))
	// Install some provided extra handler to set some request's context fields.
	// Thanks to those handler, all our logs will come with some pre-populated fields.
	m = m.Append(hlog.AccessHandler(func(r *http.Request, status, size int, duration time.Duration) {
		hlog.FromRequest(r).Info().
			Str("method", r.Method).
			Str("url", r.URL.String()).
			Int("status", status).
			Int("size", size).
			Dur("duration", duration).
			Msg("")
	}))
	m = m.Append(HeaderHandler("xForwardedFor", "X-Forwarded-For"))
	m = m.Append(hlog.RemoteAddrHandler("ip"))
	m = m.Append(hlog.UserAgentHandler("userAgent"))
	m = m.Append(hlog.RefererHandler("referer"))
	m = m.Append(hlog.RequestIDHandler("requestId", "Request-Id"))
	application.Middleware = m
	return
}

// HeaderHandler adds the request's header field as a field to the context's logger
// using fieldKey as field key and headerField as the header field
func HeaderHandler(fieldKey, headerField string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if ref := r.Header.Get(headerField); ref != "" {
				log := zerolog.Ctx(r.Context())
				log.UpdateContext(func(c zerolog.Context) zerolog.Context {
					return c.Str(fieldKey, ref)
				})
			}
			next.ServeHTTP(w, r)
		})
	}
}
