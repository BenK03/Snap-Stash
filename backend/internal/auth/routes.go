package auth

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"database/sql"
)

// check if username and password are valid then store them in the DB
func postRegister(c *gin.Context, db *sql.DB) {
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
func postLogin(c *gin.Context, db *sql.DB) {
	var req LoginReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid json"})
		return
	}

	if req.Username == "" || req.Password == "" {
		c.JSON(400, gin.H{"error": "Username and password required"})
		return
	}

	// if all is good find username and corresponding password hash from DB
	var storedHash string
	err := db.QueryRow(
		"select password_hash from Users where username = ?",
		req.Username,
	).Scan(&storedHash)

	// handle user not found
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(401, gin.H{"error": "invalid credentials"})
			return
		}

		c.JSON(500, gin.H{"error": "db error"})
		return
	}





}

