package store

import (
	"fmt"
	"os"
	"strconv"

	"gitlab.com/vyra/almercadito/almercadito-api-restful/clients/models"
	"gitlab.com/vyra/almercadito/almercadito-api-restful/environment"
)

type SpreadsheetStore struct {
	spreadsheetID   string
	spreadsheetPage string
	environment     *environment.Environment
}

func (s *SpreadsheetStore) Configuration(environment *environment.Environment) error {

	s.environment = environment
	s.spreadsheetID = os.Getenv("SPREADSHEET_ID")
	s.spreadsheetPage = os.Getenv("SPREADSHEET_PAGE")

	return nil
}

func (s *SpreadsheetStore) Get() (*[]models.Client, error) {

	fmt.Printf("[store] spreadsheetID %v\n", s.spreadsheetID)
	fmt.Printf("[store] spreadsheetPage %v\n", s.spreadsheetPage)
	readRange := s.spreadsheetPage + "!A1:I24"
	fmt.Printf("[store] readRange %v\n", readRange)

	resp, err := s.environment.Service.Spreadsheets.Values.Get(s.spreadsheetID, readRange).Do()

	if err != nil {
		fmt.Printf("[store] error %v\n", err.Error())
		return nil, err
	}

	var clients []models.Client

	if len(resp.Values) == 0 {
		fmt.Printf("[store] vacio")
		return &[]models.Client{}, nil
	}

	fmt.Printf("[store] hay %v\n", len(resp.Values))

	for _, row := range resp.Values {

		value, err := strconv.ParseInt(row[0].(string), 16, 64)
		if err != nil {
			fmt.Printf("Convert... %v", err)
			continue
		}

		clie, err := models.New(
			value,
			row[2].(string),
			"",
			"",
			"",
			"",
			"",
			"")

		if err != nil {
			fmt.Printf("%v", err)
			continue
		}

		clients = append(clients, *clie)
	}

	return &clients, nil
}

func (s *SpreadsheetStore) GetByID() (*models.Client, error) {
	return nil, nil
}
