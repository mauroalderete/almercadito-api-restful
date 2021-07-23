package store

import (
	"gitlab.com/vyra/almercadito/almercadito-api-restful/clients/models"
)

type IClientStore interface {
	Get() (*[]model.Client, error)
	GetByID() (*model.Client, error)
	Get() (*[]models.Client, error)
	GetByID() (*models.Client, error)
}
