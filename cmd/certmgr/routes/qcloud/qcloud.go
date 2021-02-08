package qcloud

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tangx/srv-lego-certmgr/cmd/certmgr/global"
	"github.com/tangx/srv-lego-certmgr/cmd/certmgr/utils"
	"github.com/tangx/srv-lego-certmgr/pkg/legox"
)

type Client interface {
	ApplyCertificate(domains ...string) (legox.Certificate, error)
}

var root *gin.RouterGroup
var lego Client
var retryChannel = make(chan string, 20)
var retryCount = make(map[string]int)

func init() {
	lego = global.LegoQcloud
	go retryApply()
}

func AppendRoute(rg *gin.RouterGroup) {
	root = rg.Group("qcloud")
	{
		root.POST("/:domain", CreateHandler)
		root.GET("/:domain", GetHandler)
	}
}

func retryApply() {
	// 错误重试队列
	for {
		domain := <-retryChannel

		err := createCertificate(domain)
		if err != nil {
			retry := retryCount[domain]
			retryCount[domain] = retry + 1
			time.Sleep(time.Duration(retry*60) * time.Second)

			// 重试次数小于 N 次，则继续， 否则退出队列
			if retry < 4 {
				retryChannel <- domain
			}
		}
		retryCount[domain] = 0

	}
}

func CreateHandler(c *gin.Context) {
	domain := c.Param("domain")

	// 后台创建证书
	go func() {
		err := createCertificate(domain)
		// 如果错误， 则重试
		if err != nil {
			retryChannel <- domain
		}
	}()

	global.CertGenerateJob[domain] = fmt.Errorf("%s 域名证书正在申请中", domain)
	c.JSON(http.StatusOK, map[string]string{"JobID": domain, "Status": "applying ..."})

}

func GetHandler(c *gin.Context) {
	domain := c.Param("domain")
	c.Redirect(http.StatusSeeOther, fmt.Sprintf("/certmgr/cert/query/%s", domain))
}

// 创建证书
func createCertificate(domain string) error {
	cert, err := lego.ApplyCertificate(domain)
	if err != nil {
		// 当出现错误时
		global.CertGenerateJob[domain] = err

		// 写入队列等待重新创建
		// retryAppyCh <- domain
		return err
	}

	utils.PushCert(domain, cert)
	delete(global.CertGenerateJob, domain)

	return nil
}
