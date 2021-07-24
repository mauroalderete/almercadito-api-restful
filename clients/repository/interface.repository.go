package repository

import (
	model "gitlab.com/vyra/almercadito/almercadito-api-restful/clients/models"
	"gitlab.com/vyra/almercadito/almercadito-api-restful/shared/hash"
)

type IClientRepository interface {
	Get() (*[]model.Client, error)
	GetByID(id hash.Hash) (*model.Client, error)
}
