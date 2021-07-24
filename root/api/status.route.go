package api

import "github.com/gin-gonic/gin"

func (r *Root) status(g *gin.Context) {
	g.String(200, "Ok")
}
