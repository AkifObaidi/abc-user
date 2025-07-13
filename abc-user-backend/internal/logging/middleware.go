package logging

import (
    "time"

    "github.com/gin-gonic/gin"
    "github.com/sirupsen/logrus"
)

// GinLogger returns a Gin middleware handler that logs requests using logrus
func GinLogger(logger *logrus.Logger) gin.HandlerFunc {
    return func(c *gin.Context) {
        start := time.Now()

        c.Next()

        latency := time.Since(start)
        status := c.Writer.Status()

        logger.WithFields(logrus.Fields{
            "method":  c.Request.Method,
            "path":    c.Request.URL.Path,
            "status":  status,
            "latency": latency,
            "client":  c.ClientIP(),
        }).Info("Handled request")
    }
}

