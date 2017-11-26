package main

import (
	"fmt"
	"os"

	"github.com/bernardigiri/restfulUserAuth/config"
	"github.com/bernardigiri/restfulUserAuth/controller"
	"github.com/bernardigiri/restfulUserAuth/model"
	"github.com/gorilla/mux"
)

func explainUsage() {
	fmt.Println("restfulServer configurationFile [encryptionKey]")
	fmt.Println("    configurationFile - Configuration file path. (required)")
	fmt.Println("    encryptionKey     - Configuration encryption key file path. (optional)")
}

func main() {
	user := model.User{}
	const password = "somePasswordValue123#.*"

	user.SetPassword(password)
	_, err1 := user.Authenticate(password)
	if err1 != nil {
		fmt.Println(err1.Error())
	}

	if len(os.Args) < 2 {
		os.Stderr.WriteString("Not enough arguments\n")
		explainUsage()
		os.Exit(1)
	}
	configFilePath := os.Args[1]
	encryptionKeyPath := ""
	if len(os.Args) > 2 {
		encryptionKeyPath = os.Args[2]
	}
	application, err := config.LoadConfig(configFilePath, encryptionKeyPath)
	if err != nil {
		keyValue := "(default)"
		if encryptionKeyPath != "" {
			keyValue = encryptionKeyPath
		}
		msg := fmt.Sprintf("Could not load configuration \"%s\" with key \"%s\" due to error: %s\n",
			configFilePath,
			keyValue,
			err.Error())
		os.Stderr.WriteString(msg)
		os.Exit(1)
	}
	router := mux.NewRouter()
	controller.ConnectLoginRoutes(router, &application)
	controller.ConnectStatusRoutes(router)
	handler := application.Middleware.Then(router)
	err = application.ListenAndServe(handler)
	if err != nil {
		application.Logger.Fatal().Err(err).Msg("Failed to connect")
	}
}
