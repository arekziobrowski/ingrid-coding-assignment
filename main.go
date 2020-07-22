package main

import (
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	initializeRoutes(router)
	return router
}

func main() {
	// test URL: http://localhost:8080/routes?src=13.388860,52.517037&dst=13.397634,52.529407&dst=13.428555,52.523219
	r := setupRouter()
	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
