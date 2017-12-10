package main

import (
	"fmt"
	"os"

	"application/config"
	"application/controller"
	"application/validation"

	"github.com/gorilla/mux"
)

func explainUsage() {
	fmt.Println("restfulServer configurationFile [encryptionKey]")
	fmt.Println("    configurationFile - Configuration file path. (required)")
	fmt.Println("    encryptionKey     - Configuration encryption key file path. (optional)")
}

func main() {
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
	validation.RegisterAll()
	router := mux.NewRouter()
	controller.ConnectLoginRoutes(router, &application)
	controller.ConnectRegisterRoutes(router, &application)
	controller.ConnectStatusRoutes(router)
	handler := application.Middleware.Then(router)
	err = application.ListenAndServe(handler)
	if err != nil {
		application.Logger.Fatal().Err(err).Msg("Failed to connect")
	}
}
