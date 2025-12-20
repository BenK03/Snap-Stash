package media

import (
	"database/sql"
	"fmt"
	snapminio "snapstash/internal/storage/minio"
	"strconv"
	"strings"
	"context"
	"time"
	"github.com/gin-gonic/gin"
	minio "github.com/minio/minio-go/v7"
)

// Pipeline HTTP → Gin → Validation → MinIO → MySQL → JSON response
func PostUpload(c *gin.Context, db *sql.DB, minioClient *snapminio.Client) {

	// get user id/check if it is valid
	userIDRaw := strings.TrimSpace(c.GetHeader("X-User-ID"))
	if userIDRaw == "" { // if user id missing
		c.JSON(400, gin.H{"error": "missing X-User-ID header"})
		return
	}

	userID, err := strconv.Atoi(userIDRaw) // convert the id to a integer
	if err != nil || userID <= 0 {         // if non integer or equal to 0 or less return error
		c.JSON(400, gin.H{"error": "invalid X-User-ID header"})
		return
	}

	// pull file
	fh, err := c.FormFile("file")
	if err != nil { // if not nil then empty so show error
		c.JSON(400, gin.H{"error": "missing file (field name must be 'file')"})
		return
	}

	// show error if file size <= 0
	if fh.Size <= 0 {
		c.JSON(400, gin.H{"error": "empty file"})
		return
	}

	// reject uploads that aren't images or videos
	mimeType := strings.TrimSpace(fh.Header.Get("Content-Type"))
	if !strings.HasPrefix(mimeType, "image/") && !strings.HasPrefix(mimeType, "video/") {
		c.JSON(415, gin.H{"error": "only image/* and video/* allowed"})
		return
	}

	// check if we can open the file
	file, err := fh.Open()
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to open uploaded file"})
		return
	}
	defer file.Close()

	// create minio path where file will be stored
	objectKey := fmt.Sprintf(
		"users/%d/%d_%s",
		userID,
		time.Now().UnixNano(),
		fh.Filename,
	)

	// upload to minio
	ctx := context.Background()

	_, err = minioClient.MC.PutObject(
		ctx,
		minioClient.Bucket,
		objectKey,
		file,
		fh.Size,
		minio.PutObjectOptions{ContentType: mimeType},
	)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to upload to storage"})
		return
	}

	// translate data in to queryable fields for the DB
	// determine media type
	mediaType := "photo"
	if strings.HasPrefix(mimeType, "video/") {
		mediaType = "video"
	}

	// insert media data into media table.
	res, err := db.Exec(
		"INSERT INTO Media (user_id, storage_path, media_type) VALUES (?, ?, ?)",
		userID,
		objectKey,
		mediaType,
	)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to insert media metadata"})
		return
	}

	// get the PK of the inserted media
	mediaID, err := res.LastInsertId()
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to read inserted media id"})
		return
	}


	c.JSON(201, gin.H{
		"message":    "uploaded",
		"media_id":   mediaID,
		"object_key": objectKey,
		"filename":   fh.Filename,
		"mime_type":  mimeType,
		"size":       fh.Size,
		"user_id":    userID,
	})

}
