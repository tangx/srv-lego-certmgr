package cert

import (
	"github.com/gin-gonic/gin"
	"github.com/tangx/srv-lego-certmgr/cmd/certmgr/utils"
	"github.com/tangx/srv-lego-certmgr/pkg/httpresponse"
)

// 显示所有存在的证书与过期时间
func ListHanlder(c *gin.Context) {
	m := utils.ListValidCerts()

	httpresponse.StatusOK(c, m)

}

func ListAllHanlder(c *gin.Context) {
	m := utils.ListAll()
	httpresponse.StatusOK(c, m)
}
