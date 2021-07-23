package repository

import (
	"gitlab.com/vyra/almercadito/almercadito-api-restful/clients/models"
	"gitlab.com/vyra/almercadito/almercadito-api-restful/clients/store"
	"gitlab.com/vyra/almercadito/almercadito-api-restful/environment"
)

type ClientRepository struct {
	clients     *[]models.Client
	environment *environment.Environment
}

func New(environment *environment.Environment) (*ClientRepository, error) {
	r := &ClientRepository{
		clients:     &[]models.Client{},
		environment: environment,
	}

	return r, nil
}

func (r *ClientRepository) Reload() error {

	var _store store.IClientStore
	_store, err := store.New()
	_store.Configuration(r.environment)

	if err != nil {
		return err
	}

	clients, err := _store.Get()

	if err != nil {
		return nil
	}

	r.clients = clients

	return nil
}

func (r *ClientRepository) Get() (*[]models.Client, error) {
	return r.clients, nil
}

func (r *ClientRepository) GetByID(id int64) (*models.Client, error) {

	for _, c := range *r.clients {
		if c.ID == id {
			return &c, nil
		}
	}

	return nil, nil
}
