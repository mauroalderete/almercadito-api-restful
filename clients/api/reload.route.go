package api

import (
	"github.com/gin-gonic/gin"
)

func (c *ClientApi) ReloadClients() func(*gin.Context) {

	return func(g *gin.Context) {

		err := c.repository.Reload()
		if err != nil {
			g.String(400, err.Error())
		}

		g.Done()
	}
}
