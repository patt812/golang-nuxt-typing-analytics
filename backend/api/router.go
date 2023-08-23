package api

import (
	"github.com/gin-gonic/gin"
	"github.com/patt812/golang-nuxt-typing-analytics/middleware"
	"github.com/patt812/golang-nuxt-typing-analytics/service"
	"gorm.io/gorm"
)

func Router(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	r.Use(middleware.SetupCORS())

	kanaAnalyticsService := service.NewKanaAnalyticsService(db)

	apiGroup := r.Group("/api")
	{
		apiGroup.GET("/analyze", func(c *gin.Context) {
			kana := c.DefaultQuery("kana", "")
			if kana == "" {
				c.JSON(400, gin.H{"error": "kana parameter is required"})
				return
			}

			patterns, err := kanaAnalyticsService.KanaToPatterns(kana)
			if err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}

			c.JSON(200, gin.H{
				"patterns": patterns,
			})
		})
	}
	return r
}
