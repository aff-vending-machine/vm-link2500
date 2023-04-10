package middleware

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

// New creates a new middleware handler
func NewLogger() fiber.Handler {
	// Return new handler
	return func(c *fiber.Ctx) error {
		start := time.Now()

		// handle request
		err := c.Next()

		code := c.Response().StatusCode()
		requestID := "<unknown>"
		if c.Locals("requestid") != nil {
			casted, ok := c.Locals("requestid").(string)
			if ok {
				requestID = casted
			}
		}

		event := log.With().
			Str("request-id", requestID).
			Str("method", c.Method()).
			Str("path", c.Path()).
			Int("header-size", len(c.Request().Header.Header())).
			Int("body-size", len(c.Request().Body())).
			Str("agent", c.Get(fiber.HeaderUserAgent)).
			Str("referer", string(c.Context().Referer())).
			Str("proto", c.Protocol()).
			Str("remote-ip", c.IP()).
			Strs("server-ip", c.IPs()).
			Int("status", code).
			Int("resp-header-size", len(c.Response().Header.Header())).
			Int("resp-body-size", len(c.Response().Body())).
			Dur("latency", time.Since(start))

		if err != nil {
			event = event.Err(err)
		}

		lg := event.Logger()

		switch {
		case code >= fiber.StatusBadRequest && code < fiber.StatusInternalServerError:
			lg.Warn().Interface("response", c.Response().Body()).Msg("http client error")

		case code >= http.StatusInternalServerError:
			lg.Error().Interface("response", c.Response().Body()).Msg("http server error")

		default:
			lg.Info().Msg("request")
		}

		return err
	}
}
