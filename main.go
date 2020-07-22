package main

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	// test URL: http://localhost:8080/routes?src=13.388860,52.517037&dst=13.397634,52.529407&dst=13.428555,52.523219
	router = gin.Default()
	initializeRoutes()
	router.Run() // listen and serve on 0.0.0.0:8080
}
