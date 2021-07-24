package server

import (
	"github.com/gin-gonic/gin"
	environment "gitlab.com/vyra/almercadito/almercadito-api-restful/environment"
)

type Server struct {
	Engine *gin.Engine
}

func New(environment *environment.Environment) (*Server, error) {

	s := gin.Default()

	serv := &Server{
		Engine: s,
	}

	return serv, nil
}
