package auth

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"database/sql"
)

// check if username and password are valid then store them in the DB
func PostRegister(c *gin.Context, db *sql.DB) {
	var req RegisterReq

	// extract JSON body and put it into req variable
	// if there are no errors then bind c to req.
	// if there are errors, return
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid json"})
		return
	}

	// check if username and password are empty
	if req.Username == "" || req.Password == "" {
		c.JSON(400, gin.H{"error": "Username and password required"})
		return
	}

	// if all is good, hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to hash password"})
		return
	}

	// now store username and password in DB
	_, err = db.Exec(
		"insert into Users (username, password_hash) values (?, ?)",
		req.Username,
		string(hashedPassword),
	)
	if err != nil {
		c.JSON(500, gin.H{"error": "db insert failed"})
		return
	}
	c.JSON(201, gin.H{"status": "ok"}) // send to frontend to proceed

}


// check if login credentials are in DB
func PostLogin(c *gin.Context, db *sql.DB) {
	var req LoginReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid json"})
		return
	}

	if req.Username == "" || req.Password == "" {
		c.JSON(400, gin.H{"error": "Username and password required"})
		return
	}

	// find username and password in DB and store it in variables
	var userID int
	var storedHash string

	err := db.QueryRow(
		"SELECT user_id, password_hash FROM Users WHERE username = ?",
		req.Username,
	).Scan(&userID, &storedHash)


	// handle user not found
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(401, gin.H{"error": "invalid credentials"})
			return
		}

		c.JSON(500, gin.H{"error": "db error"})
		return
	}

	// compare stored password_hash with the password given
	if err := bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(req.Password)); err != nil {
		c.JSON(401, gin.H{"error": "invalid credentials"})
		return
	}

	// send to frontend
	c.JSON(200, struct {
		UserID   int    `json:"user_id"`
		Username string `json:"username"`
	}{
		UserID:   userID,
		Username: req.Username,
	})

}

