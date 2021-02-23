package global

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/tangx/goutils/viperx"
	"github.com/tangx/srv-lego-certmgr/pkg/legox"
	"github.com/tangx/srv-lego-certmgr/pkg/legox/alidnsprovider"
	"github.com/tangx/srv-lego-certmgr/pkg/legox/dnspodprovider"
)

var (
	Server = gin.Default()
)

// 用于保存生成的证书，方便 GET 时快速返回。 不持久化
var CertMap = make(map[string](legox.Certificate))

// 任务队列
var (
	CertGenerateJob = make(map[string]error)
)

// Flags
var (
	DnspodEnabled bool
	AlidnsEnabled bool
)

// Providers
var Providers = map[string]legox.Provider{}

func Initial() {

	// 读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		logrus.Fatal(err.Error())
	}
	// 绑定环境变量
	viper.AutomaticEnv()

	if DnspodEnabled {
		qcloud_email := viper.GetString("ADMIN_EMAIL")
		qcloud_token := viper.GetString("DNSPOD_API_KEY")
		LegoDnspod := dnspodprovider.NewDefualtClient(qcloud_email, qcloud_token)

		Providers["dnspod"] = LegoDnspod
	}

	if AlidnsEnabled {
		alidns_accesskey := viper.GetString("ALICLOUD_ACCESS_KEY")
		alidns_secretkey := viper.GetString("ALICLOUD_SECRET_KEY")
		alidns_email := viper.GetString("ADMIN_EMAIL")
		LegoAliyun := alidnsprovider.NewDefaultClient(alidns_email, alidns_accesskey, alidns_secretkey)

		Providers["alidns"] = LegoAliyun
	}

}

func init() {
	viperx.Default()
	viperx.AddConfigPaths("$HOME/certmgr")
}
