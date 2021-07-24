package store

import (
	"fmt"
	"os"

	"gitlab.com/vyra/almercadito/almercadito-api-restful/clients/models"
	"gitlab.com/vyra/almercadito/almercadito-api-restful/environment"
	"gitlab.com/vyra/almercadito/almercadito-api-restful/shared"
	"gitlab.com/vyra/almercadito/almercadito-api-restful/shared/hash"
)

type SpreadsheetStore struct {
	spreadsheetID    string
	spreadsheetPage  string
	spreadsheetRange string
	environment      *environment.Environment
}

func New() (*SpreadsheetStore, error) {

	s := &SpreadsheetStore{}

	return s, nil
}

func (s *SpreadsheetStore) Configuration(environment *environment.Environment) error {

	s.environment = environment
	s.spreadsheetID = os.Getenv("MODULE_CLIENTS_SPREADSHEET_ID")
	s.spreadsheetPage = os.Getenv("MODULE_CLIENTS_SPREADSHEET_PAGE")
	s.spreadsheetRange = os.Getenv("MODULE_CLIENTS_SPREADSHEET_RANGE")

	return nil
}

func (s *SpreadsheetStore) Get() (*[]models.Client, error) {

	resp, err := s.environment.Service.Spreadsheets.Values.Get(s.spreadsheetID, s.spreadsheetPage+"!"+s.spreadsheetRange).Do()

	if err != nil {
		fmt.Printf("[Clients.Store.Get] Error %v\n", err.Error())
		return nil, err
	}

	var clients []models.Client

	if len(resp.Values) == 0 {
		return &[]models.Client{}, nil
	}

	for _, row := range resp.Values {

		var h hash.Hash

		h.SetFromHex(row[0].(string))

		clie, err := models.New(
			h,
			shared.GetStringFromRow(row, 2, ""),
			shared.GetStringFromRow(row, 3, ""),
			shared.GetStringFromRow(row, 4, ""),
			shared.GetStringFromRow(row, 5, ""),
			shared.GetStringFromRow(row, 6, ""),
			shared.GetStringFromRow(row, 7, ""),
			shared.GetStringFromRow(row, 8, ""))

		if err != nil {
			fmt.Printf("[Clients.Store.Get] Error to Created Model %v", err)
			continue
		}

		clients = append(clients, *clie)
	}

	return &clients, nil
}

func (s *SpreadsheetStore) GetByID(id hash.Hash) (*models.Client, error) {

	resp, err := s.environment.Service.Spreadsheets.Values.Get(s.spreadsheetID, s.spreadsheetPage+"!"+s.spreadsheetRange).Do()

	if err != nil {
		fmt.Printf("[Clients.Store.GetByID] Error %v\n", err.Error())
		return nil, err
	}

	for _, row := range resp.Values {

		if id.EqualFromHex(row[0].(string)) {

			clie, err := models.New(
				id,
				shared.GetStringFromRow(row, 2, ""),
				shared.GetStringFromRow(row, 3, ""),
				shared.GetStringFromRow(row, 4, ""),
				shared.GetStringFromRow(row, 5, ""),
				shared.GetStringFromRow(row, 6, ""),
				shared.GetStringFromRow(row, 7, ""),
				shared.GetStringFromRow(row, 8, ""))

			return clie, err
		}
	}

	return nil, nil
}
