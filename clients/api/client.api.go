package api

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/vyra/almercadito/almercadito-api-restful/clients/repository"
	environment "gitlab.com/vyra/almercadito/almercadito-api-restful/environment"
	"gitlab.com/vyra/almercadito/almercadito-api-restful/server"
)

type ClientApi struct {
	Server      *gin.Engine
	Environment *environment.Environment
	repository  *repository.ClientRepository
}

func New(srv *server.Server, environment *environment.Environment) (*ClientApi, error) {

	repo, err := repository.New(environment)
	if err != nil {
		return nil, err
	}

	clientApi := &ClientApi{
		Environment: environment,
		Server:      srv.Engine,
		repository:  repo,
	}

	return clientApi, nil
}

func (c *ClientApi) Setup() error {

	c.Server.GET("/clients", c.GetClients())
	c.Server.GET("/client/:id", c.GetClientById())

	return nil
}
