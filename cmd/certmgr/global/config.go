package global

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/tangx/srv-lego-certmgr/pkg/legox"
	"github.com/tangx/srv-lego-certmgr/pkg/legox/dnspodprovider"
)

var (
	// Server  = ginx.NewDefaultServer()
	Server  = gin.Default()
	AppName = "lego-certmgr"
)

// lego provider clients
var (
	qcloud_email = os.Getenv("DNSPOD_API_EMAIL")
	qcloud_token = os.Getenv("DNSPOD_API_KEY")
	LegoQcloud   = dnspodprovider.NewDefualtClient(qcloud_email, qcloud_token)
)

// 用于保存生成的证书，方便 GET 时快速返回。 不持久化
var CertMap = make(map[string](legox.Certificate))

// 任务队列
var (
	CertGenerateJob = make(map[string]error)
	CertGenerateCh  = make(chan string, 20)
)
