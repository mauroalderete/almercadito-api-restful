package repository

import (
	"gitlab.com/vyra/almercadito/almercadito-api-restful/clients/models"
	"gitlab.com/vyra/almercadito/almercadito-api-restful/clients/store"
	"gitlab.com/vyra/almercadito/almercadito-api-restful/environment"
	"gitlab.com/vyra/almercadito/almercadito-api-restful/shared/hash"
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

func (r *ClientRepository) Get() (*[]models.Client, error) {

	var _store store.IClientStore
	_store, err := store.New()
	_store.Configuration(r.environment)

	if err != nil {
		return nil, err
	}

	clients, err := _store.Get()

	if err != nil {
		return nil, err
	}

	r.clients = clients

	return r.clients, nil
}

func (r *ClientRepository) GetByID(id hash.Hash) (*models.Client, error) {

	var _store store.IClientStore
	_store, err := store.New()
	_store.Configuration(r.environment)

	if err != nil {
		return nil, err
	}

	client, err := _store.GetByID(id)

	if err != nil {
		return nil, err
	}

	return client, nil
}
