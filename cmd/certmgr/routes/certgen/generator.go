package certgen

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/tangx/srv-lego-certmgr/cmd/certmgr/global"
	"github.com/tangx/srv-lego-certmgr/cmd/certmgr/utils"
	"github.com/tangx/srv-lego-certmgr/pkg/httpresponse"
)

// 包环境变量设置 retry 设置
var (
	// retry 通道， 为了支持多 provider 的情况
	retryChannel = make(map[string]chan string)
)

func ApplyCertificateHandler(c *gin.Context) {

	domains := sortDomains(c.Param("domain"))

	// 如果已存在证书，且在有效期内
	// 则直接返回
	cert, ok := utils.GetCert(domains)
	if ok {
		now := time.Now().Local()
		if now.Sub(cert.NotBefore.Local()) < time.Hour*24*15 {
			logrus.Infof("证书 %s 在有效期内， 无需重建", domains)
			httpresponse.StatusDefault(c, http.StatusCreated, "domain cert created", nil)
			return
		}
	}

	// 获取 url 中 dns provider 的位置。
	prov := providerPostion(c.Request.URL.Path, 3)

	// 后台执行申请
	go func() {
		err := applyCertificate(prov, domains)
		if err != nil {
			logrus.Errorf("%s apply failed: %s ", domains, err.Error())

			// 加入重试队列
			retryChannel[prov] <- domains
			return
		}
	}()

	httpresponse.StatusDefault(c, http.StatusCreated, "domain cert created", nil)
}

func applyCertificate(prov string, domains string) error {
	dl := splitDomains(domains)

	// 申请新证书
	cert, err := global.Providers[prov].ApplyCertificate(dl...)
	if err != nil {
		return err
	}
	// 缓存结果
	utils.PushCert(domains, cert)
	return nil
}

// retryApply 错误重试
func retryApply(prov string) {
	ch := make(chan string, 20)
	retryChannel[prov] = ch

	go func() {
		logrus.Infof("启动 %s 重试队列", prov)
		for {
			domains := <-retryChannel[prov]

			for i := 1; i < 4; i++ {
				logrus.Infof("%s -> 第 %d 次重试:  %s\n", prov, i, domains)
				err := applyCertificate(prov, domains)
				if err == nil {
					break
				}

				time.Sleep(30 * time.Second)
			}
		}
	}()

}
