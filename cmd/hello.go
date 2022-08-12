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

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		logger.Debug().
			Str("method", c.Request.Method).
			Str("endpoint", c.Request.RequestURI).
			Msg("Pong")

		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
