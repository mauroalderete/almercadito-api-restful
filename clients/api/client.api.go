package api

import (
	"github.com/gin-gonic/gin"
	environment "gitlab.com/vyra/almercadito/almercadito-api-restful/environment"
	"gitlab.com/vyra/almercadito/almercadito-api-restful/server"
)

type ClientApi struct {
	Server      *gin.Engine
	Environment *environment.Environment
}

func New(srv *server.Server, environment *environment.Environment) (*ClientApi, error) {

	repo, err := repository.New(environment)
	if err != nil {
		return nil, err
	}

	clientApi := &ClientApi{
		Environment: environment,
		Server:      srv.Engine,
	}

	return clientApi, nil
}

func (c *ClientApi) Setup() error {
	c.Server.GET("/clients", c.GetClients())

	return nil
}
