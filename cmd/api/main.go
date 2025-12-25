package main

import (
	"fmt"
	"log"

	"github.com/yourname/music-shop/database"
	httpserver "github.com/yourname/music-shop/internal/http"
	"github.com/yourname/music-shop/internal/service/auth"
	"github.com/yourname/music-shop/internal/user"
)

func main() {
	db := database.InitDB()
	defer database.CloseDB(db)

	db.AutoMigrate(&user.User{})

	userRepo := user.NewGormUserRepository(db)
	authSvc := auth.NewAuthService(userRepo)

	r := httpserver.SetupRouter(authSvc)

	fmt.Println("Server running at http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
