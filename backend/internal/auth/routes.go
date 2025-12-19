package auth

import "github.com/gin-gonic/gin"


// check if username and password are valid then store them in the DB
func postRegister(c *gin.Context) {
	var req RegisterReq

	// extract JSON body and put it into req variable
	// if there are no errors then bind c to req.
	// if there are errors, return
	if err := c.ShouldBindJSON(&req); err != nil {
		return
	}

	// check if username and password are empty
	if req.Username == "" || req.Password == "" {
		c.JSON(400, gin.H{"error": "Username and password required"})
		return
	}


}


// check if login credentials are in DB
func postLogin(c *gin.Context) {
	// TODO
}

