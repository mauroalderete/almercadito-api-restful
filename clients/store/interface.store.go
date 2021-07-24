package store

import (
	"gitlab.com/vyra/almercadito/almercadito-api-restful/clients/models"
	"gitlab.com/vyra/almercadito/almercadito-api-restful/environment"
	"gitlab.com/vyra/almercadito/almercadito-api-restful/shared/hash"
)

type IClientStore interface {
	Configuration(environment *environment.Environment) error
	Get() (*[]models.Client, error)
	GetByID(id hash.Hash) (*models.Client, error)
}
