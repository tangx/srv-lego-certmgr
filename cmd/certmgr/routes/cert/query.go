package cert

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/tangx/srv-lego-certmgr/cmd/certmgr/global"
	"github.com/tangx/srv-lego-certmgr/cmd/certmgr/utils"
	"github.com/tangx/srv-lego-certmgr/pkg/httpresponse"
)

func GetHandler(c *gin.Context) {

	domain := c.Param("domain")

	cert, ok := utils.PopCert(domain)

	// 存在 则显示结果
	if ok {
		httpresponse.StatusOK(c, cert)
		return
	}

	// 不存在则判断
	err, ok := global.CertGenerateJob[domain]
	if !ok {
		// 不存在任务
		httpresponse.StatusNotFound(c, fmt.Errorf("domain certs does not exists"))
		return
	}

	// 返回当前错误
	if err != nil {
		httpresponse.StatusBadRequest(c, err)
		return
	}
}

func DownloadHandler(c *gin.Context) {
	domain := c.Param("domain")

	zipfile, err := download(domain)
	if err != nil {
		httpresponse.StatusNotFound(c, err)
		return
	}

	c.FileAttachment(zipfile, zipfile)
}
