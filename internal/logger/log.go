package log

import (
	"io"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"

	"github.com/rs/zerolog/log"
)

func NewConsoleLogger() zerolog.Logger {
	logger := log.Output(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: time.RFC3339,
	})

	return logger
}

func NewJSONLogger() zerolog.Logger {
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()

	return logger
}

func NewDiscardLogger() zerolog.Logger {
	logger := log.Output(io.Discard)
	zerolog.SetGlobalLevel(zerolog.Disabled)

	return logger
}

func GetLogger() zerolog.Logger {
	env := os.Getenv("ENV")

	env = strings.TrimSpace(env)
	env = strings.ToLower(env)

	switch env {
	case "dev", "development":
		return NewConsoleLogger()
	case "prod", "production":
		return NewJSONLogger()
	default:
		return NewDiscardLogger()
	}
}

func LoggerMiddleware(logger zerolog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		c.Next()

		latency := time.Since(t)
		status := c.Writer.Status()
		requestURI := c.Request.RequestURI
		method := c.Request.Method

		logger.Debug().
			Str("method", method).
			Int("status", status).
			Dur("latency_ms", latency).
			Str("endpoint", requestURI).
			Msg("")
	}
}
