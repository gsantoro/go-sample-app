package main

import (
	"net/http"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/gin-gonic/gin"
	"github.com/zsais/go-gin-prometheus" //nolint:goimports  // indirectly imported with `ginprometheus`
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	logger := log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	r := gin.Default()
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
	// logger.Debug().
	// 	Str("method", c.Request.Method).
	// 	Str("endpoint", c.Request.RequestURI).
	// 	Msg("Pong")

	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
