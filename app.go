package main

import (
	"backend/router"

	"github.com/gin-gonic/gin"
)

func run() {
	var httpServer *gin.Engine

	httpServer = gin.Default()

	router.Register(httpServer)

	serverAddr := "0.0.0.0:8080"

	// listen and serve on 0.0.0.0:8080
	httpServer.Run(serverAddr)
}

func main() {
	run()
}
