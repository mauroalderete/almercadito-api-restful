package main

import (
	"flag"
	"fmt"
	"log"
	"path"

	"github.com/gin-gonic/gin"
	"gitlab.com/vyra/almercadito/almercadito-api-restful/almercadito_context"
	"gitlab.com/vyra/almercadito/almercadito-api-restful/clients"
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

func Startup(context *almercadito_context.Context) *gin.Engine {

	server := gin.Default()

	server.GET("/status", func(g *gin.Context) {
		g.String(200, "Ok")
	})

	clientsApp := clients.NewClientsApp(context, server)

	clientsApp.Configure("/clients")
	clientsApp.Load()

	return server
}
