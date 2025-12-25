package http

import (
	"github.com/gin-gonic/gin"
	"github.com/yourname/music-shop/internal/http/handler"
	"github.com/yourname/music-shop/internal/http/middleware"
	"github.com/yourname/music-shop/internal/service/auth"
)

func SetupRouter(authSvc *auth.AuthService) *gin.Engine {
	r := gin.Default()

	authHandler := handler.NewAuthHandler(authSvc)

	r.POST("/auth/register", authHandler.Register)
	r.POST("/auth/login", authHandler.Login)

	api := r.Group("/api")
	api.Use(middleware.AuthRequired())
	{
		api.GET("/profile", func(c *gin.Context) {
			id := c.GetUint("user_id")
			c.JSON(200, gin.H{"user_id": id})
		})
	}

	admin := api.Group("/admin")
	admin.Use(middleware.RoleRequired("admin"))
	{
		admin.GET("/dashboard", func(c *gin.Context) {
			c.JSON(200, gin.H{"success": true})
		})
	}

	return r
}
