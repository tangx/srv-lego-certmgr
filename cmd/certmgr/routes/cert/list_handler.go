package cert

import (
	"net/http"
	"time"

	"github.com/tangx/srv-lego-certmgr/cmd/certmgr/global"
	"github.com/gin-gonic/gin"
)

// 显示所有存在的证书与过期时间
func ListHanlder(c *gin.Context) {
	m := map[string]time.Time{}

	for _, cert := range global.CertMap {
		m[cert.Domain] = cert.NotAfter
	}

	c.JSON(http.StatusOK, m)
}
