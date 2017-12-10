package controller

import (
	"net/http"

	"github.com/bernardigiri/httpServerBoilerplate/rest"
	"application/config"
	"application/dao"
	"application/handlers"
	"application/model"
	"application/responders"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"gopkg.in/mgo.v2"
)

type RegistrationForm struct {
	Login          string `valid:"alphanum"`
	Password       string `valid:"password"`
	Firstname      string `valid:"unicodeWord"`
	Lastname       string `valid:"unicodeWord"`
	Email          string `valid:"email"`
	Phonenumber    string `valid:"dialstring"`
	Enable2fa      bool   `valid:"bool"`
	SendNewsLetter bool   `valid:"bool"`
}

func ConnectRegisterRoutes(router *mux.Router, application *config.Application) (err error) {
	router.HandleFunc("/register", handlers.WithDb(application, registerEndpoint)).Methods("POST")
	return
}

func registerEndpoint(db *mgo.Database, w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		responders.ReportBadParams(w, r, err)
	}
	registrationForm := new(RegistrationForm)
	if err := schema.NewDecoder().Decode(registrationForm, r.Form); err != nil {
		responders.ReportBadParams(w, r, err)
	}
	_, exists := dao.UserByLogin(registrationForm.Login, db)
	if exists == nil {
		responders.ReportError(w, r, nil, rest.ErrorDuplicateEntry, "User exists.")
		return
	}
	user, err := model.NewUser(registrationForm.Login,
		registrationForm.Firstname,
		registrationForm.Lastname,
		registrationForm.Email,
		registrationForm.Phonenumber,
		registrationForm.Password,
		true,
		registrationForm.SendNewsLetter,
		registrationForm.Enable2fa)
	if err != nil {
		responders.ReportError(w, r, err, rest.ErrorBadParams, "Could not create user.")
		return
	}
	err = dao.UserInsert(&user, db)
	if err != nil {
		responders.ReportError(w, r, err, rest.ErrorUnknown, "Could not add user to database")
		return
	}
	responders.ReportSuccess(w, r, true)
}
