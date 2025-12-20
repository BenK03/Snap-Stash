package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// configure db
	db, err := sql.Open("mysql", "root:rootpassword@tcp(localhost:3306)/snapstash")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// configure router
	router := gin.Default()

	// run router
	router.Run("localhost:8080")
}
