package cert

import (
	"net/http"

	"github.com/tangx/srv-lego-certmgr/cmd/certmgr/global"
	"github.com/tangx/srv-lego-certmgr/cmd/certmgr/utils"
	"github.com/gin-gonic/gin"
)

func GetHandler(c *gin.Context) {

	domain := c.Param("domain")
	domain = convert(domain)

	cert, ok := utils.PopCert(domain)

	// 存在 则显示结果
	if ok {
		c.JSON(http.StatusOK, cert)
		return
	}

	// 不存在则判断
	err, ok := global.CertGenerateJob[domain]
	if !ok {
		// 不存在任务
		c.String(http.StatusBadRequest, "domain certs does not exists")
		return
	}

	// 返回当前错误
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
}

func DownloadHandler(c *gin.Context) {
	domain := c.Param("domain")
	if cert, ok := utils.PopCert(domain); ok {
		c.Redirect(http.StatusTemporaryRedirect, cert.CertStableURL)
	}
}
