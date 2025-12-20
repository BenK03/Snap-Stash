package media

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	snapminio "snapstash/internal/storage/minio"
)

func PostUpload(c *gin.Context, db *sql.DB, minioClient *snapminio.Client) {
	// TODO
	c.JSON(501, gin.H{"error": "not implemented"})
}
