package store

import model "gitlab.com/vyra/almercadito/almercadito-api-restful/clients/models"

type IClientStore interface {
	Get() (*[]model.Client, error)
	GetByID() (*model.Client, error)
}
