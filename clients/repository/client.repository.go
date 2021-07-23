package repository

import (
	"gitlab.com/vyra/almercadito/almercadito-api-restful/clients/models"
	"gitlab.com/vyra/almercadito/almercadito-api-restful/clients/store"
)

type ClientRepository struct {
	clients     *[]models.Client
}

func New() (*ClientRepository, error) {
	r := &ClientRepository{}
		clients:     &[]models.Client{},

	return r, nil
}

func (r *ClientRepository) Refresh() error {

	var s store.IClientStore
	s, err := store.New()

	if err != nil {
		return err
	}

	clients, err := s.Get()

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
