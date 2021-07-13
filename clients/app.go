package clients

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/vyra/almercadito/almercadito-api-restful/almercadito_context"
	"gitlab.com/vyra/almercadito/almercadito-api-restful/base"
)

type ClientsApp struct {
	base.App
}

type IClientsApp interface {
	Configure(route string)
	Load()
}

func NewClientsApp(context *almercadito_context.Context, server *gin.Engine) *ClientsApp {

	a := base.NewApp(context, server)

	ca := new(ClientsApp)

	ca.App = *a

	return ca
}

func (c *ClientsApp) Configure(route string) {

	var b base.App

	b.Configuration(route)

	c.Routes = append(c.Routes, base.Route{
		Route:    route,
		Context:  c.Context,
		Server:   c.Server,
		Method:   "GET",
		Function: GetRoute(c.Context),
	})
}
