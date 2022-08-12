package main

import (
	"os"

	"net/http"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/gin-gonic/gin"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	logger := log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		logger.Debug().
			Str("method", c.Request.Method).
			Str("endpoint", c.Request.RequestURI).
			Msg("Pong")

		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	err := router.SetTrustedProxies(nil)
	if err != nil {
		logger.Error().Err(err)
	}

	addr := "0.0.0.0:8080"
	err = router.Run(addr)
	if err != nil {
		logger.Error().
			Str("endpoint", addr).
			Msg("Cannot start server")
	}
}
