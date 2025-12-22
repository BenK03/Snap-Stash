package main

import (
	"database/sql"
	"log"
	"os"
	"snapstash/internal/auth"
	"snapstash/internal/config"
	"snapstash/internal/media"
	"snapstash/internal/cache"
	snapminio "snapstash/internal/storage/minio"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	cfg := config.Load()

	// configure minio client
	minioClient, err := snapminio.NewClient(cfg.MinIO)
	if err != nil {
		log.Fatalf("failed to init minio: %v", err)
	}
	_ = minioClient

	// configure redis
	rdb, err := cache.NewRedisClient()
	if err != nil {
		log.Fatalf("failed to init redis: %v", err)
	}
	_ = rdb

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

	// Endpoints
	// auth routing
	authGroup := router.Group("/api/auth")
	authGroup.POST("/register", func(c *gin.Context) {
		auth.PostRegister(c, db)
	})

	authGroup.POST("/login", func(c *gin.Context) {
		auth.PostLogin(c, db)
	})

	// media upload routing
	mediaGroup := router.Group("/api/media")
	mediaGroup.POST("/upload", func(c *gin.Context) {
		media.PostUpload(c, db, minioClient)
	})

	// list media metadata
	mediaGroup.GET("", func(c *gin.Context) {
		media.GetMedia(c, db)
	})

	// send thumbnails to frontend
	mediaGroup.GET("/:media_id/file", func(c *gin.Context) {
		media.GetMediaFile(c, db, minioClient)
	})

	// delete media
	mediaGroup.DELETE("/:media_id", func(c *gin.Context) {
		media.DeleteMedia(c, db, minioClient)
	})

	// run router
	router.Run(":8080") // Gin is running and listening on this port
}
