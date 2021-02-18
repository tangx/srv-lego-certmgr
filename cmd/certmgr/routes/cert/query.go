package cert

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tangx/srv-lego-certmgr/cmd/certmgr/global"
	"github.com/tangx/srv-lego-certmgr/cmd/certmgr/utils"
)

func GetHandler(c *gin.Context) {

	domain := c.Param("domain")

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

	zipfile, err := download(domain)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}

	c.FileAttachment(zipfile, zipfile)
}
