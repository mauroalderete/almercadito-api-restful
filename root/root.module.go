package root

import (
	"gitlab.com/vyra/almercadito/almercadito-api-restful/environment"
	"gitlab.com/vyra/almercadito/almercadito-api-restful/root/api"
	"gitlab.com/vyra/almercadito/almercadito-api-restful/server"
)

type RootModule struct {
	Api         *api.Root
	Environment *environment.Environment
}

func New(server *server.Server, environment *environment.Environment) (*RootModule, error) {

	a, err := api.New(server)
	if err != nil {
		return nil, err
	}

	rootModule := &RootModule{
		Api:         a,
		Environment: environment,
	}

	return rootModule, nil
}
