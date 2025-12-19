package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"snapstash/internal/auth"
)

func main() {
	router := gin.Default()

	api := router.Group("/api")
    authGroup := api.Group("/auth")

	db, _ := sql.Open("mysql", "user:password@tcp(localhost:3306)/snapstash")
	authHandler := &auth.Handler{DB: db}
	auth.RegisterRoutes(authGroup, authHandler)

	router.Run("localhost:8080")
}
