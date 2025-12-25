package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/yourname/music-shop/internal/database"
	"github.com/yourname/music-shop/internal/user"
)

func main() {
	db := database.InitDB()
	defer database.CloseDB(db)

	db.AutoMigrate(&user.User{})

	r := gin.Default()
	
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	fmt.Println("Server running at http://localhost:8080")
	r.Run(":8080")
}
