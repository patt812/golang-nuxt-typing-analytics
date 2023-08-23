package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
)

func SetupCORS() gin.HandlerFunc {
	allowedOrigin := os.Getenv("ALLOWED_ORIGIN")

	c := cors.New(cors.Options{
		AllowedOrigins: []string{allowedOrigin},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowedHeaders: []string{"*"},
	})

	return func(context *gin.Context) {
		c.HandlerFunc(context.Writer, context.Request)

		// handle preflight request
		if context.Request.Method == http.MethodOptions {
			context.AbortWithStatus(http.StatusNoContent)
			return
		}

		context.Next()
	}
}
