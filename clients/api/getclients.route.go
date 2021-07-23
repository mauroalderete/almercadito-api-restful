package api

import (
	"github.com/gin-gonic/gin"
)

func (a *ClientApi) GetClients() func(*gin.Context) {

	return func(g *gin.Context) {

		g.String(200, "get clients")
		// var spreadsheet_id = "1BPGEDtDsiHKNfJylUFfEy9esnYY1If6SAKHW82psthA"
		// var spreadsheet_page = "Clientes"

		// readRange := spreadsheet_page + "!A1:I24"

		// resp, err := context.Service.Spreadsheets.Values.Get(spreadsheet_id, readRange).Do()

		// if err != nil {
		// 	g.String(400, err.Error())
		// 	return
		// }

		// if len(resp.Values) == 0 {
		// 	g.String(200, "{}")
		// 	return
		// } else {
		// 	var clients []model.Client
		// 	for _, row := range resp.Values {

		// 		value, err := strconv.ParseInt(row[0].(string), 16, 64)
		// 		if err != nil {
		// 			fmt.Printf("Convert... %v", err)
		// 			continue
		// 		}

		// 		client, err := NewClient(
		// 			value,
		// 			row[2].(string),
		// 			"",
		// 			"",
		// 			"",
		// 			"",
		// 			"",
		// 			"")

		// 		if err != nil {
		// 			fmt.Printf("%v", err)
		// 			continue
		// 		}

		// 		clients = append(clients, *client)
		// 	}
		// 	g.JSON(200, clients)
		// }
	}
}
