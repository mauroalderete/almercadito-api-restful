package main

import (
	"flag"
	"fmt"
	"log"
	"path"

	"gitlab.com/vyra/almercadito/almercadito-api-restful/clients"
	"gitlab.com/vyra/almercadito/almercadito-api-restful/environment"
	"gitlab.com/vyra/almercadito/almercadito-api-restful/root"
	"gitlab.com/vyra/almercadito/almercadito-api-restful/server"
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

	var env *environment.Environment = &environment.Environment{}

	err := env.Initialize(config.CredentialFile, config.TokenFile)
	if err != nil {
		log.Fatalf("[Main] Load environment error. %v", err)
	}

	srv, err := server.New(env)
	if err != nil {
		log.Fatalf("[Main] Create server error. %v", err)
	}

	rootModule, err := root.New(srv, env)
	if err != nil {
		log.Fatalf("[Main] Error to load root module. %v", err)
	}

	err = rootModule.Api.Setup()
	if err != nil {
		log.Fatalf("[Main] Error to setup api root. %v", err)
	}

	clientsModule, err := clients.New(srv, env)
	if err != nil {
		log.Fatalf("[Main] Error to load clients module. %v", err)
	}

	err = clientsModule.Api.Setup()
	if err != nil {
		log.Fatalf("[Main] Error to setup api clients. %v", err)
	}

	srv.Engine.Run(":8080")
}
