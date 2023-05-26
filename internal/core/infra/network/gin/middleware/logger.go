package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// New creates a new middleware handler
func NewLogger() gin.HandlerFunc {
	// Return new handler
	return func(c *gin.Context) {
		start := time.Now()

		// handle request
		c.Next()

		statusCode := c.Writer.Status()
		if statusCode == 0 {
			statusCode = http.StatusOK
		}

		event := log.With().
			Str("method", c.Request.Method).
			Str("url", c.Request.URL.String()).
			Str("path", c.FullPath()).
			Str("agent", c.Request.UserAgent()).
			Str("referer", c.Request.Referer()).
			Str("proto", c.Request.Proto).
			Str("remote-ip", c.RemoteIP()).
			Str("client-ip", c.ClientIP()).
			Int("status", statusCode).
			Dur("latency", time.Since(start))

		lg := event.Logger()

		switch {
		case statusCode >= http.StatusBadRequest && statusCode < http.StatusInternalServerError:
			lg.Warn().Msg("http client error")

		case statusCode >= http.StatusInternalServerError:
			lg.Error().Msg("http server error")

		default:
			lg.Info().Msg("request")
		}
	}
}
