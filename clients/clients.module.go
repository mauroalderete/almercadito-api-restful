package clients

import (
	"gitlab.com/vyra/almercadito/almercadito-api-restful/clients/api"
	"gitlab.com/vyra/almercadito/almercadito-api-restful/clients/repository"
	environment "gitlab.com/vyra/almercadito/almercadito-api-restful/environment"
	"gitlab.com/vyra/almercadito/almercadito-api-restful/server"
)

type ClientsModule struct {
	Repository *repository.ClientRepository
	Api        *api.ClientApi
}

func New(server *server.Server, environment *environment.Environment) (*ClientsModule, error) {

	r, err := repository.New()
	if err != nil {
		return nil, err
	}

	a, err := api.New(server, environment)
	if err != nil {
		return nil, err
	}

	cm := &ClientsModule{
		Repository: r,
		Api:        a,
	}

	return cm, nil
}
