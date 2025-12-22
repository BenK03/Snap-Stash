package media

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	snapminio "snapstash/internal/storage/minio"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	minio "github.com/minio/minio-go/v7"
	"github.com/redis/go-redis/v9"
)

// Pipeline HTTP → Gin → Validation → MinIO → MySQL → JSON response
func PostUpload(c *gin.Context, db *sql.DB, minioClient *snapminio.Client) {

	// validate user ID
	userID, err := VerifyUserID(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
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
		// calls MySQL and Inserts userID, objectKey,mediaType and returns res and err
		"INSERT INTO Media (user_id, object_key, media_type) VALUES (?, ?, ?)",
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

func GetMedia(c *gin.Context, db *sql.DB) {
	// validate user ID
	userID, err := VerifyUserID(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	rows, err := db.Query(
		// calls MySQL and selects media table where user_id matches the given userID
		`SELECT media_id, object_key, media_type, created_at
		FROM Media
		WHERE user_id = ?
		ORDER BY created_at DESC, media_id DESC`,
		userID,
	)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to query media"})
		return
	}
	defer rows.Close()

	items := []MediaItem{}

	// iterate through each row we got from query and turn them into MediaItem structs
	for rows.Next() {
		var item MediaItem

		err := rows.Scan(
			&item.MediaID,
			&item.ObjectKey,
			&item.MediaType,
			&item.CreatedAt,
		)
		if err != nil {
			c.JSON(500, gin.H{"error": "failed to read media row"})
			return
		}

		items = append(items, item)
	}

	// send response
	c.JSON(200, gin.H{
		"items": items,
	})
}

// Stream to client
// Pipeline HTTP → Validation → Redis → MySQL → MinIO → Client
func GetMediaFile(c *gin.Context, db *sql.DB, minioClient *snapminio.Client, rdb *redis.Client) {

	// validate user ID
	userID, err := VerifyUserID(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// get media id from URL param and validate it
	mediaID, err := VerifyMediaID(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// redis cache  (demo defaults)
	cacheKey := MediaBytesCacheKey(mediaID)
	cacheCTypeKey := MediaCTypeCacheKey(mediaID)

	cacheTTL := 60 * time.Second            // cache expires after 60 seconds
	cacheMaxBytes := int64(5 * 1024 * 1024) // 5MB max guard

	// try redis cache first
	ctx := context.Background()

	cachedBytes, err := rdb.Get(ctx, cacheKey).Bytes()
	if err == nil {
		cachedCType := strings.TrimSpace(rdb.Get(ctx, cacheCTypeKey).Val())
		if cachedCType == "" {
			cachedCType = "application/octet-stream"
		}

		// cache hit: send cached bytes
		c.Data(200, cachedCType, cachedBytes)
		return
	}

	// cache miss: look up object_key for this media
	var objectKey string

	err = db.QueryRow(
		`SELECT object_key
		FROM Media
		WHERE media_id = ? AND user_id = ?
		LIMIT 1`,
		mediaID,
		userID,
	).Scan(&objectKey)

	if err == sql.ErrNoRows {
		c.JSON(404, gin.H{"error": "media not found"})
		return
	}
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to look up media"})
		return
	}

	// if all is good fetch the media from minio and stream to client
	ctxStorage := context.Background()

	obj, err := minioClient.MC.GetObject(
		ctxStorage,
		minioClient.Bucket,
		objectKey,
		minio.GetObjectOptions{},
	)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to read from storage"})
		return
	}
	defer obj.Close()

	stat, err := obj.Stat()
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to stat stored object"})
		return
	}

	contentType := stat.ContentType
	if strings.TrimSpace(contentType) == "" {
		contentType = "application/octet-stream"
	}

	// store frequently viewed media in redis cache
	// if small enough, cache in redis
	if CanCacheMedia(stat.Size, cacheMaxBytes) {
		data, err := io.ReadAll(obj)
		if err != nil {
			c.JSON(500, gin.H{"error": "failed to read stored object"})
			return
		}

		// store in bytes & content-type in redis
		_ = rdb.Set(ctx, cacheKey, data, cacheTTL).Err()
		_ = rdb.Set(ctx, cacheCTypeKey, contentType, cacheTTL).Err()

		c.Data(200, contentType, data)
		return
	}

	// send from minio
	c.DataFromReader(
		200,
		stat.Size,
		contentType,
		obj,
		nil,
	)

}

// Delete selected media
func DeleteMedia(c *gin.Context, db *sql.DB, minioClient *snapminio.Client) {
	// validate user ID
	userID, err := VerifyUserID(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	mediaID, err := VerifyMediaID(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// look up object_key for this media and enforce ownership
	var objectKey string

	err = db.QueryRow(
		`SELECT object_key
		FROM Media
		WHERE media_id = ? AND user_id = ?
		LIMIT 1`,
		mediaID,
		userID,
	).Scan(&objectKey)

	if err == sql.ErrNoRows {
		c.JSON(404, gin.H{"error": "media not found"})
		return
	}
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to look up media"})
		return
	}

	// delete object from minio
	ctx := context.Background()

	err = minioClient.MC.RemoveObject(
		ctx,
		minioClient.Bucket,
		objectKey,
		minio.RemoveObjectOptions{},
	)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to delete from storage"})
		return
	}

	// delete row from database
	_, err = db.Exec(
		`DELETE FROM Media
		WHERE media_id = ? AND user_id = ?`,
		mediaID,
		userID,
	)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to delete media metadata"})
		return
	}

	// send http status
	c.Status(204)

}
