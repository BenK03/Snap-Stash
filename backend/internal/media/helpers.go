package media

import (
	"errors"
	"strconv"
	"strings"
	"github.com/gin-gonic/gin"
)

func VerifyUserID(c *gin.Context) (int, error) {
	userIDRaw := strings.TrimSpace(c.GetHeader("X-User-ID"))
	if userIDRaw == "" {
		return 0, errors.New("missing X-User-ID header")
	}

	userID, err := strconv.Atoi(userIDRaw)
	if err != nil || userID <= 0 {
		return 0, errors.New("invalid X-User-ID header")
	}

	return userID, nil
}