package main

import (
	"snapstash/internal/config"
	"snapstash/internal/auth"
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	cfg := config.Load()
	_ = cfg

	// configure db
	db, err := sql.Open("mysql", "root:rootpassword@tcp(localhost:3306)/snapstash")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// configure router
	router := gin.Default()

	// routing
	authGroup := router.Group("/api/auth")
	authGroup.POST("/register", func(c *gin.Context) {
		auth.PostRegister(c, db)
	})

	authGroup.POST("/login", func(c *gin.Context) {
		auth.PostLogin(c, db)
	})

	// run router
	router.Run("localhost:8080") // Gin is running and listening on this port
}
