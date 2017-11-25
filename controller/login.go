package controller

import (
	"net/http"

	"github.com/bernardigiri/restfulUserAuth/config"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/hlog"
)

func ConnectLoginRoutes(router *mux.Router, application config.Application) (err error) {
	router.HandleFunc("/login", loginEndpoint).Methods("GET")
	return
}

func loginEndpoint(w http.ResponseWriter, req *http.Request) {
	hlog.FromRequest(req).
		Warn().
		Str("user", "current user").
		Str("status", "ok").
		Msg("Something happened")
	w.Write([]byte("Hello Go!"))
}
