package controller

import (
	"net/http"

	"github.com/bernardigiri/restfulUserAuth/responders"
	"github.com/gorilla/mux"
)

func ConnectStatusRoutes(router *mux.Router) (err error) {
	router.HandleFunc("/status", statusEndPoint).Methods("GET")
	router.HandleFunc("/hello", helloEndPoint).Methods("GET")
	return
}

func statusEndPoint(w http.ResponseWriter, r *http.Request) {
	responders.ReportSuccess(w, r, true)
}

func helloEndPoint(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello."));
}
