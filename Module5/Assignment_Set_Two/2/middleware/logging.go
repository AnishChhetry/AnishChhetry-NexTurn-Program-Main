package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		c.Next()

		duration := time.Since(startTime)
		fmt.Printf("Request Method: %s | Path: %s | Status: %d | Duration: %v\n", c.Request.Method, c.Request.URL.Path, c.Writer.Status(), duration)
	}
}
