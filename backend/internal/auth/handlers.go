package auth

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"database/sql"
)

type Handler struct{
	DB *sql.DB
}

type registerReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type loginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// accepts JSON body with username and password
func (h *Handler) Register(c *gin.Context) {
	var req registerReq
	_ = c.ShouldBindJSON(&req)

	// Check if username and password are provided
	if req.Username == "" || req.Password == "" {
		c.JSON(400, gin.H{"error": "username and password required"})
		return
	}

	// temporary
	if h.DB == nil {
		c.JSON(500, gin.H{"error": "db not initialized"})
		return
	}

	// Hashes password using bcrypt
	hash, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	_ = hash

	_, _ = h.DB.Exec(
		"INSERT INTO Users (username, password_hash) VALUES (?, ?)",
		req.Username,
		string(hash),
	)

	// I will make DB calls later (use this for now)
	c.JSON(201, gin.H{"status": "ok"})

}

func (h *Handler) Login(c *gin.Context) {
	var req loginReq
	_ = c.ShouldBindJSON(&req)

	// Check if username and password are provided
	if req.Username == "" || req.Password == "" {
		c.JSON(400, gin.H{"error": "username and password required"})
		return
	}

}