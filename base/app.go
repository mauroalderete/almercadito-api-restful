package base

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/vyra/almercadito/almercadito-api-restful/almercadito_context"
)

type App struct {
	Context *almercadito_context.Context
	Server  *gin.Engine
	Routes  []Route
}

type IApp interface {
	Configuration(route string)
	Load()
}

func NewApp(context *almercadito_context.Context, server *gin.Engine) *App {

	ca := new(App)

	ca.Context = context
	ca.Server = server

	return ca
}

func (a *App) Configuration(route string) {
}

func (a *App) Load() {
	for _, route := range a.Routes {
		route.Load()
	}
}
