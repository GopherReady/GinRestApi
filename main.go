package main

import (
	"log"
	"net/http"

	"github.com/GopherReady/GinRestApi/router"
	"github.com/gin-gonic/gin"
)

func main() {
	// Create the Gin engine
	g := gin.New()

	// Gin middlewares
	var middleware []gin.HandlerFunc

	// Routes
	router.Load(
		// Cores.
		g,

		// middlewares
		middleware...,
	)
	log.Printf("Start to listening the incoming requests on http address 127.0.0.1%s", ":8080")
	log.Printf(http.ListenAndServe(":8080", g).Error())
}
