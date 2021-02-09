package cert

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tangx/srv-lego-certmgr/cmd/certmgr/global"
)

// 显示所有存在的证书与过期时间
func ListHanlder(c *gin.Context) {
	m := make(map[string]time.Time)
	for _, cert := range global.CertMap {
		m[cert.Domain] = cert.NotAfter
	}

	c.JSON(http.StatusOK, m)
}
