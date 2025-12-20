package main

import (
	"snapstash/internal/config"
	"snapstash/internal/auth"
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	snapminio "snapstash/internal/storage/minio"
	"log"
	"os"
)

func main() {
	cfg := config.Load()

	// configure minio client
	minioClient, err := snapminio.NewClient(cfg.MinIO)
	if err != nil {
		log.Fatalf("failed to init minio: %v", err)
	}
	_ = minioClient

	// configure db
	dsn := os.Getenv("MYSQL_DSN")
	if dsn == "" {
		dsn = "root:rootpassword@tcp(localhost:3306)/snapstash"
	}
	db, err := sql.Open("mysql", dsn)

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
	router.Run(":8080") // Gin is running and listening on this port
}
