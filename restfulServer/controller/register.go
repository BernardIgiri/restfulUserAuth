package controller

import (
	"net/http"

	"github.com/bernardigiri/httpServerBoilerplate/rest"
	"github.com/bernardigiri/restfulUserAuth/config"
	"github.com/bernardigiri/restfulUserAuth/dao"
	"github.com/bernardigiri/restfulUserAuth/handlers"
	"github.com/bernardigiri/restfulUserAuth/model"
	"github.com/bernardigiri/restfulUserAuth/responders"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
)

func ConnectRegisterRoutes(router *mux.Router, application *config.Application) (err error) {
	router.HandleFunc("/register", handlers.WithDb(application, registerEndpoint)).Methods("POST")
	return
}

func registerEndpoint(db *mgo.Database, w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var (
		login          = r.Form.Get("login")
		password       = r.Form.Get("password")
		firstname      = r.Form.Get("firstname")
		lastname       = r.Form.Get("lastname")
		email          = r.Form.Get("email")
		phonenumber    = r.Form.Get("phone")
		enable2fa      = r.Form.Get("enable2fa") == "true"
		sendNewsLetter = r.Form.Get("sendNewsLetter") == "true"
	)
	_, exists := dao.UserByLogin(login, db)
	if exists == nil {
		responders.ReportError(w, r, nil, rest.ErrorDuplicateEntry, "User exists.")
		return
	}
	user, err := model.NewUser(login, firstname, lastname, email, phonenumber, password, true, sendNewsLetter, enable2fa)
	if err != nil {
		responders.ReportError(w, r, err, rest.ErrorBadParams, "Could not create user.")
		return
	}
	responders.ReportSuccess(w, r, true)
	err = dao.UserInsert(&user, db)
	if err != nil {
		responders.ReportError(w, r, err, rest.ErrorUnknown, "Could not add user to database")
	}
}
