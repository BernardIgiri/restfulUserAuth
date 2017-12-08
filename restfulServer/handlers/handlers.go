package handlers

import (
	"net/http"

	"github.com/alexedwards/scs"
	"github.com/bernardigiri/httpServerBoilerplate/rest"
	"github.com/bernardigiri/restfulUserAuth/config"
	"github.com/bernardigiri/restfulUserAuth/dao"
	"github.com/bernardigiri/restfulUserAuth/model"
	"github.com/bernardigiri/restfulUserAuth/responders"
	"gopkg.in/mgo.v2"
)

// Controls application session
type SessionModel interface {
	StartSession(w http.ResponseWriter, r *http.Request, user *model.User) (session *scs.Session, err error)
}

// Modifies an http handler to take a database connection
func WithDb(application *config.Application, fn func(*mgo.Database, http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		connection, err := application.GetDatabaseConnection()
		defer connection.Close()
		if err != nil {
			responders.Report500(w, r, err, "Could not connect to database")
			return
		}
		fn(connection.DB(""), w, r)
	}
}

// Modifies an http handler to take a session model and a database connection
func WithSessionModel(application *config.Application, fn func(SessionModel, *mgo.Database, http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		connection, err := application.GetDatabaseConnection()
		defer connection.Close()
		if err != nil {
			responders.Report500(w, r, err, "Could not connect to database")
			return
		}
		fn(application, connection.DB(""), w, r)
	}
}

// Modifies an http handler to take a session and a database connection
func WithDbSession(application *config.Application, fn func(*scs.Session, *mgo.Database, http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		connection, err := application.GetDatabaseConnection()
		defer connection.Close()
		if err != nil {
			responders.Report500(w, r, err, "Could not connect to database")
			return
		}
		session := application.GetSession(r)
		fn(session, connection.DB(""), w, r)
	}
}

// Modifies an http handler to take an authenticated session and a database connection
func WithRequiredLogin(application *config.Application, fn func(*model.User, *scs.Session, *mgo.Database, http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		connection, err := application.GetDatabaseConnection()
		defer connection.Close()
		if err != nil {
			responders.Report500(w, r, err, "Could not connect to database")
			return
		}
		session := application.GetSession(r)
		login, err := session.GetString(model.UserLoginSessionKey)
		if err != nil {
			responders.ReportError(w, r, err, rest.ErrorNotAuthenticated, "User not attached to session")
			return
		}
		user, err := dao.UserByLogin(login, connection.DB(""))
		if err != nil {
			responders.ReportError(w, r, err, rest.ErrorNotAuthenticated, "Authenticated user not found")
			return
		}
		fn(user, session, connection.DB(""), w, r)
	}
}
