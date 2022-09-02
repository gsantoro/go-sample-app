package main

import (
	"fmt"
	"os"

	_ "github.com/gsantoro/go-sample-app/docs"
	log "github.com/gsantoro/go-sample-app/internal/logger"
	"github.com/gsantoro/go-sample-app/internal/rest"

	"github.com/gin-gonic/gin"
	ginprometheus "github.com/zsais/go-gin-prometheus"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Go Sample App
// @version         1.0
// @description     Sample application in Go
// @contact.name   	Giuseppe Santoro
// @license.name  	Apache 2.0
// @license.url   	http://www.apache.org/licenses/LICENSE-2.0.html
// @host      		localhost:8080
// @BasePath  		/api/v1
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

	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping", rest.Ping)
		v1.GET("/health", rest.Health)

		// http://localhost:8080/api/v1/docs/index.html
		v1.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

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
