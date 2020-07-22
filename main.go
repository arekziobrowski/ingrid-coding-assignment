package main

import (
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	initializeRoutes(router)
	return router
}

// @title Ingrid Backend Coding Task App
// @version 1.0
// @description This is a sample REST API application built for Ingrid coding task.

// @contact.name Arkadiusz Ziobrowski
// @contact.email arekziobrowski@gmail.com
// @query.collection.format multi

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	// test URL: http://localhost:8080/routes?src=13.388860,52.517037&dst=13.397634,52.529407&dst=13.428555,52.523219
	r := setupRouter()
	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
