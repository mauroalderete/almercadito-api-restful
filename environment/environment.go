package environment

import (
	"context"
	"log"

	"gitlab.com/rayquen-google/golang/auth/auth_service"
	"gitlab.com/rayquen-google/golang/auth/auth_service_spreadsheet"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

type Environment struct {
	Auth    auth_service.IAuthService
	Service *sheets.Service
}

type IEnvironment interface {
	Initialize(credentialFile string, tokenFile string) error
}

func (e *Environment) Initialize(credentialFile string, tokenFile string) error {
	e.Auth = &auth_service_spreadsheet.AuthServiceSpreadsheet{}

	err := e.Auth.Initialize(credentialFile, tokenFile, true)
	if err != nil {
		log.Fatalf("[Main] Error al inicializar %v", err)
	}

	err = e.Auth.Authenticate()
	if err != nil {
		log.Fatalf("[Main] Error al autenticar %v", err)
	}

	srv, err := sheets.NewService(context.Background(), option.WithHTTPClient(e.Auth.GetClient()))
	if err != nil {
		log.Fatalf("[Main::NewService] %v", err)
	}

	e.Service = srv

	return nil
}
