package api

import (
	"github.com/gin-gonic/gin"
)

func (c *ClientApi) GetClients() func(*gin.Context) {

	return func(g *gin.Context) {

		clients, err := c.repository.Get()
		if err != nil {
			g.String(400, err.Error())
		}

		g.JSON(200, clients)
	}
}
