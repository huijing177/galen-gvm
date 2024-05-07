package api

import (
	"net/http"

	"galen-gvm/api/base"
	"galen-gvm/api/example"
	"galen-gvm/api/system"
	"galen-gvm/internal/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// /system/api/v1/createapi
func Router() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.IPLimit)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/ping", Ping)

	baseRouter := router.Group("base")
	base.BaseRouter(baseRouter)
	systemRouter := router.Group("system")
	systemRouter.Use(middleware.Cors, middleware.Record, middleware.JwtAuth)
	system.SystemRouter(systemRouter)
	exampleRouter := router.Group("example")
	exampleRouter.Use(middleware.Cors, middleware.Record, middleware.JwtAuth)
	example.ExampleRouter(exampleRouter)

	return router
}

// Ping godoc
//
//	@Summary		pingSummary
//	@Description	pingDescription
//	@Tags			pingTags
//	@Accept			json
//	@Produce		json
//	@Success		204	{object}	interface{}
//	@Router			/ping [get]
func Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "pong")
}
