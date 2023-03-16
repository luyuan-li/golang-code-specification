package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowMethods: []string{"GET", "POST", "PUT", "HEAD", "DELETE", "PATCH"},
		AllowHeaders: []string{"Origin", "Content-Length", "Content-Type",
			"X-Real-Ip", "X-Appengine-Remote-Addr", "Access-Control-Allow-Origin"},
		AllowCredentials: false,
		AllowAllOrigins:  true,
		MaxAge:           12 * time.Hour,
	})
}
