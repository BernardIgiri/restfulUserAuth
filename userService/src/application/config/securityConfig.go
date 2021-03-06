package config

import (
	"crypto/rand"
	"io"
	"net/http"
	"time"

	"github.com/alexedwards/scs"
	"application/model"
)

// Configuration constants
const (
	keyLength = 32
)

func loadSecurityConfig(application *Application, config Config) (err error) {
	//application.Middleware = application.Middleware.Append(nosurf.NewPure)
	sessionKey := make([]byte, keyLength)
	_, err = io.ReadFull(rand.Reader, sessionKey)
	if err != nil {
		return
	}
	application.sessionMan = scs.NewCookieManager(string(sessionKey))
	application.sessionMan.Lifetime(time.Hour)
	application.sessionMan.Persist(false)
	application.sessionMan.Secure(true)
	application.sessionMan.HttpOnly(true)
	return
}

// GetSession returns the session from the request
func (application *Application) GetSession(r *http.Request) *scs.Session {
	return application.sessionMan.Load(r)
}

// StartSession creates session for given user
func (application *Application) StartSession(w http.ResponseWriter, r *http.Request, user *model.User) (session *scs.Session, err error) {
	session = application.sessionMan.Load(r)
	err = session.PutString(w, model.UserLoginSessionKey, user.Login)
	return
}
