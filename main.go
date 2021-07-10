package main

import (
	"context"
	"fmt"
	"log"

	"gitlab.com/rayquen-google/golang/auth/auth_service"
	"gitlab.com/rayquen-google/golang/auth/auth_service_spreadsheet"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func main() {
	fmt.Println("alMercadito API RESTful")

	var auth auth_service.IAuthService = &auth_service_spreadsheet.AuthServiceSpreadsheet{}

	err := auth.Initialize("credential.json", "token.json", true)
	if err != nil {
		log.Fatalf("[Main] Error al inicializar %v", err)
	}

	err = auth.Authenticate()
	if err != nil {
		log.Fatalf("[Main] Error al autenticar %v", err)
	}

	srv, err := sheets.NewService(context.Background(), option.WithHTTPClient(auth.GetClient()))
	if err != nil {
		log.Fatalf("[Main::NewService] %v", err)
	}

	var spreadsheet_id = "1BPGEDtDsiHKNfJylUFfEy9esnYY1If6SAKHW82psthA"
	var spreadsheet_page = "Clientes"

	readRange := spreadsheet_page + "!A1:I24"

	resp, err := srv.Spreadsheets.Values.Get(spreadsheet_id, readRange).Do()

	if err != nil {
		log.Fatalf("[Main::GetValues] %v", err)
	}

	if len(resp.Values) == 0 {
		fmt.Println("Nada!!")
	} else {
		for _, row := range resp.Values {
			fmt.Printf("%s: %s\n", row[0], row[2])
		}
	}

	fmt.Println("Ok")
}

//importar gin
