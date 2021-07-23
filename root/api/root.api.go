package api

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/vyra/almercadito/almercadito-api-restful/server"
)

type Root struct {
	Server *gin.Engine
}

func New(srv *server.Server) (*Root, error) {
	r := &Root{
		Server: srv.Engine,
	}

	return r, nil
}

func (r *Root) Setup() error {

	r.Server.GET("/status", r.status)

	return nil
}
