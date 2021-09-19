package service

import "github.com/gin-gonic/gin"

func Serve(initializeServer func(r *gin.Engine), port string) {
	r := gin.Default()
	initializeServer(r)
	r.Run(port)
}
