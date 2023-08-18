package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Router(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	apiGroup := r.Group("/api")
	{
		apiGroup.GET("/analyze", func(c *gin.Context) {
			var version string
			row := db.Raw("SELECT VERSION()").Row()
			row.Scan(&version)

			c.JSON(200, gin.H{
				"version": version,
			})
		})
	}
	return r
}
