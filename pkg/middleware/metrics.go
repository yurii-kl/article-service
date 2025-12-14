package middleware

import (
	"strconv"
	"time"

	"github.com/Article/article-service/pkg/metrics"
	"github.com/gin-gonic/gin"
)

func PrometheusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.FullPath()
		if path == "" {
			path = c.Request.URL.Path
		}

		c.Next()

		duration := time.Since(start).Seconds()
		status := strconv.Itoa(c.Writer.Status())
		method := c.Request.Method

		metrics.HTTPRequestsTotal.WithLabelValues(method, path, status).Inc()
		metrics.HTTPRequestDuration.WithLabelValues(method, path).Observe(duration)

		if c.Request.ContentLength > 0 {
			metrics.HTTPRequestSize.WithLabelValues(method, path).Observe(float64(c.Request.ContentLength))
		}
	}
}
