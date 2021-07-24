package api

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/vyra/almercadito/almercadito-api-restful/shared/hash"
)

func (c *ClientApi) GetClientById() func(*gin.Context) {

	return func(g *gin.Context) {

		var h hash.Hash
		h.SetFromHex(g.Param("id"))

		client, err := c.repository.GetByID(h)
		if err != nil {
			g.String(400, err.Error())
		}

		g.JSON(200, client)
	}
}
