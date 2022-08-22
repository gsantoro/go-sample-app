package main

import (
	"fmt"
	"os"

	log "github.com/gsantoro/go-sample-app/internal/logger"
	"github.com/gsantoro/go-sample-app/internal/rest"

	"github.com/gin-gonic/gin"
	ginprometheus "github.com/zsais/go-gin-prometheus"
)

func main() {
	logger := log.GetLogger()

	// r := gin.Default()
	r := gin.New()
	r.Use(log.LoggerMiddleware(logger))

	p := ginprometheus.NewPrometheus("gin")
	p.Use(r)

	err := r.SetTrustedProxies(nil)
	if err != nil {
		logger.Error().Err(err)
	}

	r.GET("/ping", rest.Ping)
	r.GET("/health", rest.Health)

	hostname := os.Getenv("HOSTNAME")
	port := os.Getenv("PORT")

	addr := fmt.Sprintf("%s:%s", hostname, port)
	err = r.Run(addr)
	if err != nil {
		logger.Error().
			Str("endpoint", addr).
			Msg("Cannot start server")
	}
}
