package store

import (
	"gitlab.com/vyra/almercadito/almercadito-api-restful/clients/models"
	"gitlab.com/vyra/almercadito/almercadito-api-restful/environment"
)

type IClientStore interface {
	Configuration(environment *environment.Environment) error
	Get() (*[]models.Client, error)
	GetByID() (*models.Client, error)
}

func New() (*SpreadsheetStore, error) {

	s := &SpreadsheetStore{}

	return s, nil
}
