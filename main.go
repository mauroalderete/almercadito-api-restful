package main

import (
	"flag"
	"fmt"
	"log"
	"path"

	"github.com/gin-gonic/gin"
	"gitlab.com/vyra/almercadito/almercadito-api-restful/almercadito_context"
)

type Config struct {
	CredentialFile string
	TokenFile      string
}

func main() {
	fmt.Println("alMercadito API RESTful")

	flagWorkdir := flag.String("workdir", "./", "Working directory. For default is './'")
	flagCredential := flag.String("credential", "credential.json", "Credential filename. For default is 'credential.json' ")
	flagToken := flag.String("token", "token.json", "Token filename. For default is 'token.json'")
	flagVersion := flag.Bool("version", false, "Version of alMercadito API RESTful Login")
	flagHelp := flag.Bool("help", false, "Show help information")

	flag.Parse()

	if *flagVersion {
		fmt.Println("v1.0.0")
		return
	}

	if *flagHelp {
		flag.PrintDefaults()
		return
	}

	var config Config = Config{
		CredentialFile: path.Join(*flagWorkdir, *flagCredential),
		TokenFile:      path.Join(*flagWorkdir, *flagToken),
	}

	ctx, err := Configuration(&config)
	if err != nil {
		log.Fatalf("[Main] Hubo un problema al configurar el servicio %v", err)
	}

	Startup(ctx).Run()
}

func Configuration(config *Config) (*almercadito_context.Context, error) {
	var ctx *almercadito_context.Context = &almercadito_context.Context{}

	err := ctx.Initialize(config.CredentialFile, config.TokenFile)

	if err != nil {
		return nil, err
	} else {
		return ctx, nil
	}
}

type Client struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func Startup(ctx *almercadito_context.Context) *gin.Engine {

	server := gin.Default()

	server.GET("/status", func(g *gin.Context) {
		g.String(200, "Ok")
	})

	server.GET("/clients", func(g *gin.Context) {
		var spreadsheet_id = "1BPGEDtDsiHKNfJylUFfEy9esnYY1If6SAKHW82psthA"
		var spreadsheet_page = "Clientes"

		readRange := spreadsheet_page + "!A1:I24"

		resp, err := ctx.Service.Spreadsheets.Values.Get(spreadsheet_id, readRange).Do()

		if err != nil {
			g.String(400, err.Error())
			return
		}

		if len(resp.Values) == 0 {
			g.String(200, "{}")
			return
		} else {
			var clients []Client
			for _, row := range resp.Values {
				//fmt.Printf("%s: %s\n", row[0], row[2])

				clients = append(clients, Client{
					ID:   row[0].(string),
					Name: row[2].(string),
				})
			}
			g.JSON(200, clients)
		}
	})

	return server
}
