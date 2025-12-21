package media

type UploadResponse struct {
	MediaID   int64  `json:"media_id"`
	ObjectKey string `json:"object_key"`
	Filename  string `json:"filename"`
	MimeType  string `json:"mime_type"`
	Size      int64  `json:"size"`
}

type MediaItem struct {
	MediaID   int64  `json:"media_id"`
	ObjectKey string `json:"object_key"`
	MediaType string `json:"media_type"`
	CreatedAt string `json:"created_at"`
}