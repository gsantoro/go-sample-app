package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Health        Used to check the health of the web api
// @Summary      Health
// @Description  Responds with status=ok.
// @Tags         Health
// @Produce      json
// @Router       /health [get]
func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}
