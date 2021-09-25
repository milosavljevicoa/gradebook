package main

import (
	"milosavljevicoa/gradebook/app/registry"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.POST("/services", registry.AddService)
	r.DELETE("/services", registry.RemoveService)
	r.Run(registry.ServerPort)
}
