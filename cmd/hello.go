package main

import (
	"net/http"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/gin-gonic/gin"
	ginprometheus "github.com/zsais/go-gin-prometheus"
)

func main() {
	// zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs
	logger := log.Output(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: time.RFC3339,
	})

	// r := gin.Default()
	r := gin.New()
	r.Use(Logger(logger))
	r.GET("/ping", Ping)

	p := ginprometheus.NewPrometheus("gin")
	p.Use(r)

	err := r.SetTrustedProxies(nil)
	if err != nil {
		logger.Error().Err(err)
	}

	addr := "0.0.0.0:8080"
	err = r.Run(addr)
	if err != nil {
		logger.Error().
			Str("endpoint", addr).
			Msg("Cannot start server")
	}
}

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func Logger(logger zerolog.Logger) gin.HandlerFunc {
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
