package store

import model "gitlab.com/vyra/almercadito/almercadito-api-restful/clients/models"

type SpreadsheetStore struct {
}

func New() (*SpreadsheetStore, error) {

	s := &SpreadsheetStore{}

	return s, nil
}

func (s *SpreadsheetStore) Get() (*[]model.Client, error) {
	return nil, nil
}

func (s *SpreadsheetStore) GetByID() (*model.Client, error) {
	return nil, nil
}
