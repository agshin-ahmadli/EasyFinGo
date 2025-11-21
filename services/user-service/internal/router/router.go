package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	r.GET("/health", healthCheck)
	
	v1 := r.Group("/api/v1")
	{
		_ = v1
	}
}

func healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  "ok",
		"message": "User service is running",
		"service": "user-service",
	})
}
