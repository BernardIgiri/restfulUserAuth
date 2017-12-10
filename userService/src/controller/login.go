package controller

import (
	"net/http"

	"github.com/alexedwards/scs"
	"github.com/bernardigiri/httpServerBoilerplate/rest"
	"application/config"
	"application/dao"
	"application/handlers"
	"application/model"
	"application/responders"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
)

func ConnectLoginRoutes(router *mux.Router, application *config.Application) (err error) {
	router.HandleFunc("/login", handlers.WithSessionModel(application, loginEndpoint)).Methods("POST")
	router.HandleFunc("/logout", handlers.WithRequiredLogin(application, logoutEndpoint)).Methods("POST")
	return
}

func loginEndpoint(sessionModel handlers.SessionModel, db *mgo.Database, w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var (
		login    = r.Form.Get("login")
		password = r.Form.Get("password")
	)
	user, err := dao.UserByLogin(login, db)
	if err != nil {
		responders.ReportDebugError(w, r, err, rest.ErrorNotAuthorized, "Could not find login")
		return
	}
	match, err := user.Authenticate(password)
	if err != nil {
		responders.ReportError(w, r, err, rest.ErrorNotAuthorized, "Error logging in")
		return
	}
	if !match {
		responders.ReportDebugError(w, r, err, rest.ErrorNotAuthorized, "Invalid password")
		return
	}
	sessionModel.StartSession(w, r, user)
	responders.ReportSuccess(w, r, true)
}

func logoutEndpoint(user *model.User, session *scs.Session, db *mgo.Database, w http.ResponseWriter, r *http.Request) {
	err := session.Destroy(w)
	if err != nil {
		responders.ReportError(w, r, err, rest.ErrorInvalidSession, "Invalid session")
		return
	}
	responders.ReportSuccess(w, r, true)
}
