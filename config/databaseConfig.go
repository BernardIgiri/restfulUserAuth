package config

import (
	"github.com/bernardigiri/restfulUserAuth/encryption"
	"gopkg.in/mgo.v2"
)

func (a *Application) GetDatabaseConnection() (database *mgo.Database, err error) {
	session, err := mgo.Dial(a.db.server)
	session.Login(&a.db.credential)
	database = session.DB(a.db.database)
	return
}

func loadDatabaseConfig(application *Application, config Config, decrypter encryption.Decrypter) (err error) {
	application.db.server = config.Db.Server
	password, err := decrypter.Decrypt(config.Db.Password)
	application.db.credential = mgo.Credential{Username: config.Db.Username, Password: string(password)}
	return
}
