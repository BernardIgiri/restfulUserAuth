package config

import (
	"fmt"
	"strconv"
	"time"

	"github.com/bernardigiri/restfulUserAuth/encryption"
	"gopkg.in/mgo.v2"
)

var session *mgo.Session

func (a *Application) GetDatabaseConnection() (database *mgo.Database, err error) {
	if session == nil {
		user := string(a.db.username)
		password := string(a.db.password)
		mongoDBDialInfo := &mgo.DialInfo{
			Addrs:    []string{a.db.server + ":" + strconv.Itoa(a.db.port)},
			Timeout:  60 * time.Second,
			Database: a.db.database,
			Username: user,
			Password: password,
		}
		fmt.Printf("%+v\n", mongoDBDialInfo)
		session, err = mgo.DialWithInfo(mongoDBDialInfo)
		if err != nil {
			return
		}
		session.SetMode(mgo.Monotonic, true)
	}
	database = session.DB(a.db.database)
	return
}

func loadDatabaseConfig(application *Application, config Config, decrypter encryption.Decrypter) (err error) {
	password, err := decrypter.Decrypt(config.Db.Password)
	application.db.port = config.Db.Port
	application.db.server = config.Db.Server
	application.db.database = config.Db.Database
	application.db.username = []byte(config.Db.Username)
	application.db.password = []byte(password)
	return
}
