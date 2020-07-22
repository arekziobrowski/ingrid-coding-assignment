package main

import "ingrid-coding-assignment/route"

func initializeRoutes() {
	router.GET("/routes", route.GetRoutes)
}
