package main

import (
	"github.com/gin-gonic/gin"
	"snapstash/internal/auth"
)

func main() {
	router := gin.Default()

	api := router.Group("/api")
    authGroup := api.Group("/auth")

	authHandler := &auth.Handler{}
	auth.RegisterRoutes(authGroup, authHandler)

	router.Run("localhost:8080")
}
