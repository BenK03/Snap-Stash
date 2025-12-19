package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)
func pingHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "pong"})
}

func main() {
	router := gin.Default()
	router.GET("/ping", pingHandler)
	router.Run("localhost:8080")
}