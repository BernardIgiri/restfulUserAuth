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

type SessionManager interface {
	StartSession(w http.ResponseWriter, r *http.Request, user *model.User) (session *scs.Session, err error)
}

func WithDb(application *config.Application, fn func(*mgo.Database, http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		db, err := application.GetDatabaseConnection()
		if err != nil {
			responders.Report500(w, r, err, "Could not connect to database")
			return
		}
		fn(db, w, r)
	}
}

func WithSessionManager(application *config.Application, fn func(SessionManager, *mgo.Database, http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		db, err := application.GetDatabaseConnection()
		if err != nil {
			responders.Report500(w, r, err, "Could not connect to database")
			return
		}
		fn(application, db, w, r)
	}
}

func WithDbSession(application *config.Application, fn func(*scs.Session, *mgo.Database, http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		db, err := application.GetDatabaseConnection()
		if err != nil {
			responders.Report500(w, r, err, "Could not connect to database")
			return
		}
		session := application.GetSession(r)
		fn(session, db, w, r)
	}
}

func WithRequiredLogin(application *config.Application, fn func(*model.User, *scs.Session, *mgo.Database, http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		db, err := application.GetDatabaseConnection()
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
		user, err := dao.UserByLogin(login, db)
		if err != nil {
			responders.ReportError(w, r, err, rest.ErrorNotAuthenticated, "Authenticated user not found")
			return
		}
		fn(user, session, db, w, r)
	}
}
