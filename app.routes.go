package main

import (
	"github.com/gin-gonic/gin"
	"ingrid-coding-assignment/route"
)

func initializeRoutes(router *gin.Engine) {
	router.GET("/routes", route.GetRoutes)
}
