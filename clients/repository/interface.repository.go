package repository

import model "gitlab.com/vyra/almercadito/almercadito-api-restful/clients/models"

type IClientRepository interface {
	Refresh() error
	Get() (*[]model.Client, error)
	GetByID(id int64) (*model.Client, error)
}
