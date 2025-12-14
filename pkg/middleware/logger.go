package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

const (
	RequestIDHeader = "X-Request-ID"
	RequestIDKey    = "request_id"
	LoggerKey       = "logger"
)

func RequestLoggerMiddleware(logger *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path

		requestID := c.GetHeader(RequestIDHeader)
		if requestID == "" || !isValidUUID(requestID) {
			requestID = uuid.New().String()
		}

		c.Set(RequestIDKey, requestID)
		c.Header(RequestIDHeader, requestID)

		requestLogger := logger.WithFields(logrus.Fields{
			"request_id": requestID,
			"method":     c.Request.Method,
			"path":       path,
			"ip":         c.ClientIP(),
		})

		c.Set(LoggerKey, requestLogger)

		requestLogger.Info("request started")

		c.Next()

		latency := time.Since(start)
		statusCode := c.Writer.Status()

		logFields := logrus.Fields{
			"status":     statusCode,
			"latency_ms": latency.Milliseconds(),
			"size_bytes": c.Writer.Size(),
		}

		if query := c.Request.URL.RawQuery; query != "" {
			logFields["query"] = query
		}

		if len(c.Errors) > 0 {
			logFields["errors"] = c.Errors.String()
		}

		logEntry := requestLogger.WithFields(logFields)
		switch {
		case statusCode >= 500:
			logEntry.Error("request completed")
		case statusCode >= 400:
			logEntry.Warn("request completed")
		case latency > 3*time.Second:
			logEntry.Warn("slow request completed")
		default:
			logEntry.Info("request completed")
		}
	}
}

func isValidUUID(s string) bool {
	_, err := uuid.Parse(s)
	return err == nil
}
