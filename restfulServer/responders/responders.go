package responders

import (
	"net/http"

	"github.com/bernardigiri/httpServerBoilerplate/rest"
	"github.com/rs/zerolog/hlog"
)

func Report500(w http.ResponseWriter, r *http.Request, err error, message string) {
	w.WriteHeader(http.StatusInternalServerError)
	hlog.FromRequest(r).
		Fatal().
		Err(err).
		Msg(message)
}

func ReportBrokenPipe(r *http.Request, err error) {
	hlog.FromRequest(r).
		Fatal().
		Err(err).
		Msg("Could not write to response")
}

func ReportBadParams(w http.ResponseWriter, r *http.Request, errorEvent error) {
	w.WriteHeader(http.StatusBadRequest)
	msg := rest.ErrorBadParams
	ReportError(w, r, errorEvent, msg, msg)
}

func ReportDebugError(w http.ResponseWriter, r *http.Request, errorEvent error, errorName, message string) {
	hlog.FromRequest(r).
		Debug().
		Err(errorEvent).
		Msg(message)
	if err := rest.WriteJSONError(w, rest.ErrorNotAuthorized); err != nil {
		ReportBrokenPipe(r, err)
	}
}

func ReportError(w http.ResponseWriter, r *http.Request, errorEvent error, errorName, message string) {
	hlog.FromRequest(r).
		Error().
		Err(errorEvent).
		Msg(message)
	if err := rest.WriteJSONError(w, errorName); err != nil {
		ReportBrokenPipe(r, err)
	}
}

func ReportSuccess(w http.ResponseWriter, r *http.Request, object interface{}) {
	if err := rest.WriteJSONSuccess(w, object); err != nil {
		ReportBrokenPipe(r, err)
	}
}
