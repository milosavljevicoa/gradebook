package main

import (
	"milosavljevicoa/gradebook/app/registry"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/services", registry.ServeHttpPost)
	r.Run(registry.ServerPort)
}
