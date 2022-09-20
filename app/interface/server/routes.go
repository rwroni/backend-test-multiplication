package server

import (
	"github.com/gin-gonic/gin"
	"github.com/gomodul/envy"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"project-name/app/transport"
)

func setupRouter(transport *transport.Tp, app *gin.Engine) {
	// allow swagger
	if envy.Get("ALLOW_SWAGGER", "false") == "true" {
		app.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	// this router group for readiness and health check
	ping := app.Group("/ping")
	ping.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	v1 := app.Group("/v1")
	sample := v1.Group("/multiplication")
	sample.GET("/", transport.SampleTransport.MultiplicationTable)

}
