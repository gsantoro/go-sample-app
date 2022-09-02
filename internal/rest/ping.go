package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping          Used to check if the web server is up
// @Summary      Ping
// @Description  Responds with Pong.
// @Tags         ping
// @Produce      json
// @Router       /ping [get]
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
