package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"project-name/app/interface/container"
	"project-name/app/transport"
)

func StartServer(container container.Container) {
	// init gin
	app := gin.New()

	// setup transport and router (parent)
	transport := transport.SetupTransport(&container)
	setupRouter(transport, app)

	// run gin apps
	fmt.Println(app.Run())
}
