package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// /system/api/v1/createapi
func Router() *gin.Engine {
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	systemRouter := router.Group("system")
	SystemRouter(systemRouter)
	exampleRouter := router.Group("example")
	ExampleRouter(exampleRouter)

	return router
}

func SystemRouter(g *gin.RouterGroup) {
	apiGroup := g.Group("api")
	apiGroupV1 := apiGroup.Group("v1")
	apiGroupV1.GET("/ping", Ping)
}

func ExampleRouter(g *gin.RouterGroup) {

}

// ShowAccount godoc
// @Summary      Show an account
// @Description  ping func
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Router       /accounts/{id} [get]
func Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "pong")
}
