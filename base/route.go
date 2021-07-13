package base

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/vyra/almercadito/almercadito-api-restful/almercadito_context"
)

type Route struct {
	Context  *almercadito_context.Context
	Server   *gin.Engine
	Route    string
	Method   string
	Function func(*gin.Context)
}

type IRoute interface {
	Initialize(ctx *almercadito_context.Context, server *gin.Engine) *gin.Engine
	Configuration(route string)
	Load()
}

func NewRoute(context *almercadito_context.Context, server *gin.Engine) *Route {

	r := new(Route)
	r.Context = context
	r.Server = server
	r.Route = ""
	r.Function = nil

	return r
}

func (r *Route) Configuration(route string) {
}

func (r *Route) Load() {

	if r.Method == "GET" {
		r.Server.GET(r.Route, r.Function)
	}
}
