package config

import (
	"strconv"
	"time"

	"application/encryption"

	"gopkg.in/mgo.v2"
)

var session *mgo.Session

func (a *Application) GetDatabaseConnection() (connection *mgo.Session, err error) {
	connection = session.Copy()
	return
}

func loadDatabaseConfig(application *Application, config Config, decrypter encryption.Decrypter) (err error) {
	password, err := decrypter.Decrypt(config.Db.Password)
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{config.Db.Server + ":" + strconv.Itoa(config.Db.Port)},
		Timeout:  2 * time.Minute,
		Database: config.Db.Database,
		Username: config.Db.Username,
		Password: string(password),
	}
	session, err = mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		return
	}
	session.SetMode(mgo.Monotonic, true)
	return
}
