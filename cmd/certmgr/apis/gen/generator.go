package gen

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-jarvis/rum-gonic/pkg/httpx"
	"github.com/sirupsen/logrus"
	"github.com/tangx/srv-lego-certmgr/cmd/certmgr/utils"
	"github.com/tangx/srv-lego-certmgr/pkg/httpresponse"
)

type GenerateCertByDomain struct {
	httpx.MethodPost `path:"/:provider/:domain"`
	Provider         string `uri:"provider"`
	Domain           string `uri:"domain"`
}

func (req *GenerateCertByDomain) Output(c *gin.Context) (interface{}, error) {
	// 如果已存在证书，且在有效期内
	// 则直接返回
	cert, ok := utils.GetCert(req.Domain)
	if ok {
		now := time.Now().Local()
		if now.Sub(cert.NotBefore.Local()) < time.Hour*24*15 {
			logrus.Infof("证书 %s 在有效期内， 无需重建", req.Domain)
			ret := httpresponse.RespDefault(http.StatusCreated, "domain cert created", nil)
			return ret, nil
		}
	}

	// 后台执行申请
	// go func() {
	// 	err := applyCertificate(req.Provider, req.Domain)
	// 	if err != nil {
	// 		logrus.Errorf("%s apply failed: %s ", req.Domain, err.Error())

	// 		// 加入重试队列
	// 		retryChannel[req.Provider] <- req.Domain
	// 		return
	// 	}
	// }()

	ret := httpresponse.RespDefault(http.StatusCreated, "domain cert created", nil)
	return ret, nil
}

// func applyCertificate(prov string, domains string) error {
// 	dl := splitDomains(domains)

// 	// 申请新证书
// 	cert, err := global.OldProviders[prov].ApplyCertificate(dl...)
// 	if err != nil {
// 		return err
// 	}
// 	// 缓存结果
// 	utils.PushCert(domains, cert)
// 	return nil
// }
