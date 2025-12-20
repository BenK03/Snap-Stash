package media

import (
	"database/sql"
	"strconv"
	"strings"
	"github.com/gin-gonic/gin"
	snapminio "snapstash/internal/storage/minio"
)

func PostUpload(c *gin.Context, db *sql.DB, minioClient *snapminio.Client) {
	_ = db
	_ = minioClient

	// get user id
	userIDRaw := strings.TrimSpace(c.GetHeader("X-User-ID"))
	if userIDRaw == "" { // if user id missing
		c.JSON(400, gin.H{"error": "missing X-User-ID header"})
		return
	}

	userID, err := strconv.Atoi(userIDRaw) // convert the id to a integer
	if err != nil || userID <= 0 { // if non integer or equal to 0 or less return error
		c.JSON(400, gin.H{"error": "invalid X-User-ID header"})
		return
	}



}
