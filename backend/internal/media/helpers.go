package media

import (
	"errors"
	"strconv"
	"strings"
	"github.com/gin-gonic/gin"
)

func VerifyUserID(c *gin.Context) (int, error) {
	// get user id/check if it is valid
	userIDRaw := strings.TrimSpace(c.GetHeader("X-User-ID"))
	if userIDRaw == "" { // if user id missing
		return 0, errors.New("missing X-User-ID header")
	}

	userID, err := strconv.Atoi(userIDRaw) // convert the id to a integer
	if err != nil || userID <= 0 { // if non integer or equal to 0 or less return error
		return 0, errors.New("invalid X-User-ID header")
	}

	return userID, nil
}