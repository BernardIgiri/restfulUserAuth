package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/alexedwards/scs"
	"github.com/bernardigiri/httpServerBoilerplate/rest"
	"github.com/bernardigiri/restfulUserAuth/config"
	"github.com/bernardigiri/restfulUserAuth/dao"
	"github.com/bernardigiri/restfulUserAuth/handlers"
	"github.com/bernardigiri/restfulUserAuth/model"
	"github.com/bernardigiri/restfulUserAuth/responders"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
)

func ConnectLoginRoutes(router *mux.Router, application *config.Application) (err error) {
	router.HandleFunc("/login", handlers.WithSessionManager(application, loginEndpoint)).Methods("POST")
	router.HandleFunc("/logout", handlers.WithRequiredLogin(application, logoutEndpoint)).Methods("POST")
	return
}

func getUserFromRequest(r *http.Request) (user *model.User, err error) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return
	}
	err = json.Unmarshal(body, user)
	return
}

func loginEndpoint(sMan handlers.SessionManager, db *mgo.Database, w http.ResponseWriter, r *http.Request) {
	userLogin, err := getUserFromRequest(r)
	if err != nil {
		responders.ReportBadParams(w, r, err)
		return
	}
	user, err := dao.UserByLogin(userLogin.Login, db)
	if err != nil {
		responders.ReportDebugError(w, r, err, rest.ErrorNotAuthorized, "Could not find login")
		return
	}
	match, err := user.Authenticate(userLogin.Password)
	if err != nil {
		responders.ReportError(w, r, err, rest.ErrorNotAuthorized, "Error logging in")
		return
	}
	if !match {
		responders.ReportDebugError(w, r, err, rest.ErrorNotAuthorized, "Invalid password")
		return
	}
	sMan.StartSession(w, r, user)
	responders.ReportSuccess(w, r, true)
}

func logoutEndpoint(user *model.User, session *scs.Session, db *mgo.Database, w http.ResponseWriter, r *http.Request) {
	err := session.Destroy(w)
	if err != nil {
		responders.ReportError(w, r, err, rest.ErrorInvalidSession, "Invalid session")
	}
	responders.ReportSuccess(w, r, true)
}
