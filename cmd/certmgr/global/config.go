package global

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/tangx/srv-lego-certmgr/pkg/legox"
	"github.com/tangx/srv-lego-certmgr/pkg/legox/alidnsprovider"
	"github.com/tangx/srv-lego-certmgr/pkg/legox/dnspodprovider"
)

var (
	Server  = gin.Default()
	AppName = "lego-certmgr"
)

// 用于保存生成的证书，方便 GET 时快速返回。 不持久化
var CertMap = make(map[string](legox.Certificate))

// 任务队列
var (
	CertGenerateJob = make(map[string]error)
	CertGenerateCh  = make(chan string, 20)
)

// Flags
var (
	DnspodEnabled bool
	AlidnsEnabled bool
)

// Providers
var Providers = map[string]legox.Provider{}

func Initial() {
	if DnspodEnabled {
		qcloud_email := os.Getenv("ADMIN_EMAIL")
		qcloud_token := os.Getenv("DNSPOD_API_KEY")
		LegoDnspod := dnspodprovider.NewDefualtClient(qcloud_email, qcloud_token)

		Providers["dnspod"] = LegoDnspod
	}

	if AlidnsEnabled {
		alidns_accesskey := os.Getenv("ALICLOUD_ACCESS_KEY")
		alidns_secretkey := os.Getenv("ALICLOUD_SECRET_KEY")
		alidns_email := os.Getenv("ADMIN_EMAIL")
		LegoAliyun := alidnsprovider.NewDefaultClient(alidns_email, alidns_accesskey, alidns_secretkey)

		Providers["alidns"] = LegoAliyun
	}
}
