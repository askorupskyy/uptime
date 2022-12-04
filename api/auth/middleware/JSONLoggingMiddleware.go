package middleware

import (
	"time"

	gin "github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func JSONLogMiddleware(log *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		// Process Request
		c.Next()
		// Stop timer
		duration := time.Now().Sub(start)

		log.Sugar().Named("main logger").Infow("Request",
			"duration", duration,
			"method", c.Request.Method,
			"path", c.Request.RequestURI,
			"status", c.Writer.Status(),
			"referrer", c.Request.Referer(),
			"request_id", c.Writer.Header().Get("Request-Id"),
		)

		if c.Writer.Status() >= 500 {
			log.Sugar().Named("main logger").Errorw("Request failed!")
		}
	}
}
