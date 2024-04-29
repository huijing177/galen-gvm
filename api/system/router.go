package system

import (
	"github.com/gin-gonic/gin"
)

func SystemRouter(g *gin.RouterGroup) {
	apiGroup := g.Group("api")
	ApiRouter(apiGroup)
}
