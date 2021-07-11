package almercadito_context

import (
	"context"
	"log"

	"gitlab.com/rayquen-google/golang/auth/auth_service"
	"gitlab.com/rayquen-google/golang/auth/auth_service_spreadsheet"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

type Context struct {
	Auth    auth_service.IAuthService
	Service *sheets.Service
	Self    *Context
}

type IContext interface {
	Initialize(credentialFile string, tokenFile string) error
}

func (c *Context) Initialize(credentialFile string, tokenFile string) error {
	c.Auth = &auth_service_spreadsheet.AuthServiceSpreadsheet{}

	err := c.Auth.Initialize(credentialFile, tokenFile, true)
	if err != nil {
		log.Fatalf("[Main] Error al inicializar %v", err)
	}

	err = c.Auth.Authenticate()
	if err != nil {
		log.Fatalf("[Main] Error al autenticar %v", err)
	}

	srv, err := sheets.NewService(context.Background(), option.WithHTTPClient(c.Auth.GetClient()))
	if err != nil {
		log.Fatalf("[Main::NewService] %v", err)
	}

	c.Service = srv

	return nil
}
