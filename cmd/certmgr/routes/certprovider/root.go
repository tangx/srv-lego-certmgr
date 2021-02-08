package certprovider

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/tangx/srv-lego-certmgr/cmd/certmgr/global"
	"github.com/tangx/srv-lego-certmgr/cmd/certmgr/utils"
)

// 包环境变量设置 retry 设置
var (
	// retry 计数器
	retryCounter = make(map[string]int)
	// retry 通道， 为了支持多 provider 的情况
	retryChannel = make(map[string]chan string)
)

func init() {
	retryApply()
}

// 设置路由
func AppendRoute(rg *gin.RouterGroup, name string) {

	subroot := rg.Group("gen")

	provroot := subroot.Group(name)
	// provroot.GET("/ping")
	provroot.POST("/:domains", ApplyCertificateHandler)
	provroot.GET("/:domains", GetHandler)

}

func ApplyCertificateHandler(c *gin.Context) {

	domains := SortDomains(c.Param("domains"))

	// prov for provider
	prov := ProviderPostion(c.Request.URL.Path, 3)

	// 后台执行
	go func() {
		err := applyCertificate(prov, domains)
		if err != nil {
			logrus.Errorf("%s apply failed: %s ", domains, err.Error())

			// 加入重试队列
			retryChannel[prov] <- domains
			return
		}
	}()

	c.String(http.StatusOK, "applying...")
}

// GetHandler 303 redirect
func GetHandler(c *gin.Context) {
	domains := SortDomains(c.Param("domains"))
	c.Redirect(http.StatusSeeOther, fmt.Sprintf("/certmgr/cert/query/%s", domains))
}

func applyCertificate(prov string, domains string) error {
	dl := SplitDomains(domains)

	logrus.Debugln("provider = ", prov)
	cert, err := global.Providers[prov].ApplyCertificate(dl...)

	if err != nil {
		return err
	}

	// 缓存结果
	utils.PushCert(domains, cert)
	return nil
}

// retryApply 错误重试
func retryApply() {
	retry := func(prov string) {
		for {
			domains := <-retryChannel[prov]
			times := retryCounter[domains]

			// 超过4次重试就略过
			if times > 4 {
				retryCounter[domains] = 0
				continue
			}

			// 重试
			// 等待 60 秒继续
			time.Sleep(60 * time.Second)
			err := applyCertificate(prov, domains)

			// 失败继续重试
			if err != nil {
				retryChannel[prov] <- domains
				continue
			}

			// 成功，重试次数重置为 0
			retryCounter[domains] = 0
		}
	}

	for prov := range retryChannel {
		go retry(prov)
	}
}
