package config

import (
	"errors"
	"io/ioutil"

	"application/encryption"

	"github.com/alexedwards/scs"
	"github.com/justinas/alice"
	"github.com/rs/zerolog"
	"gopkg.in/yaml.v2"
)

type Application struct {
	sessionMan *scs.Manager
	Middleware alice.Chain
	Logger     zerolog.Logger
	http       struct {
		hostname string
		port     int
	}
}

type Config struct {
	Db struct {
		Server   string
		Port     int
		Database string
		Username string
		Password string
	}
	Log struct {
		Path  string
		Level string
	}
	Http struct {
		Hostname string
		Port     int
	}
}

func LoadConfig(configPath, keyPath string) (application Application, err error) {
	application = Application{}
	config := Config{}
	application.Middleware = alice.New()
	if len(keyPath) < 1 {
		keyPath = configPath + ".key"
	}
	configData, err := ioutil.ReadFile(configPath)
	if err != nil {
		err = errors.New("config file error: " + err.Error())
		return
	}
	key, err := ioutil.ReadFile(keyPath)
	if err != nil {
		err = errors.New("key file error: " + err.Error())
		return
	}
	err = yaml.Unmarshal([]byte(configData), &config)
	if err != nil {
		return
	}
	decrypter := encryption.NewConfigDecrypter(key)
	err = loadDatabaseConfig(&application, config, decrypter)
	if err != nil {
		err = errors.New("database config error: " + err.Error())
		return
	}
	err = loadLoggerConfig(&application, config)
	if err != nil {
		err = errors.New("logger config error: " + err.Error())
		return
	}
	err = loadSecurityConfig(&application, config)
	if err != nil {
		err = errors.New("security config error: " + err.Error())
		return
	}
	err = loadHttpConfig(&application, config)
	return
}
