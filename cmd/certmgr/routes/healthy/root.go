package healthy

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AppendRoute(rg *gin.RouterGroup) {
	rg.GET("/healthy", HealthyHandler)
}

func HealthyHandler(c *gin.Context) {
	c.String(http.StatusOK, "success")
}
