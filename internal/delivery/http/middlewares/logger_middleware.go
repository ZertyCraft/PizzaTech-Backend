package middlewares

import (
    "github.com/gin-gonic/gin"
    "github.com/sirupsen/logrus"
    "time"
)

func Logger(log *logrus.Logger) gin.HandlerFunc {
    return func(c *gin.Context) {
        start := time.Now()
        c.Next()
        log.WithFields(logrus.Fields{
            "status":   c.Writer.Status(),
            "method":   c.Request.Method,
            "path":     c.Request.URL.Path,
            "duration": time.Since(start),
        }).Info("request completed")
    }
}
