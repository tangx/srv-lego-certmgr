package global

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/tangx/goutils/ginx"
	"github.com/tangx/goutils/viperx"
	"github.com/tangx/srv-lego-certmgr/pkg/container"
	"github.com/tangx/srv-lego-certmgr/pkg/legox"
	"github.com/tangx/srv-lego-certmgr/pkg/legox/alidnsprovider"
	"github.com/tangx/srv-lego-certmgr/pkg/legox/dnspodprovider"
	"github.com/tangx/srv-lego-certmgr/pkg/storage"
	"github.com/tangx/srv-lego-certmgr/pkg/storage/filesystem"
)

var (
	Server  = ginx.Default()
	Appname = "lego-certmgr"
)

var DPmapping = container.NewDomainProviderMap()

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
var Storager storage.Storager

func Initial() {
	// InitialProvider()
	InitialStorager()
}

// initial backend storage
func InitialStorager() {
	viper.AutomaticEnv()
	// BackendStorageClass
	// class := viper.GetString("BACKEND_StorageClass")
	class := viper.GetString("BACKEND_STORAGE_CLASS")
	switch class {
	case "filesystem":
		//todo
		logrus.Error("create filesystem")
		dir := viper.GetString("BACKEND_FILE_SYSTEM_DIR")
		Storager = filesystem.NewStorager(dir)
	}

	if Storager == nil {
		panic("storage is nil")
	}
}

// initial dns provider
func InitialProvider() {

	// 读取配置文件
	configHome := fmt.Sprintf("$HOME/%s", Appname)
	_ = viperx.ReadInConfig(configHome)
	// 绑定环境变量
	viper.AutomaticEnv()

	if DnspodEnabled {
		qcloud_email := viper.GetString("ADMIN_EMAIL")
		qcloud_token := viper.GetString("DNSPOD_API_KEY")
		LegoDnspod := dnspodprovider.NewDefualtClient(
			qcloud_email,
			qcloud_token)

		Providers["dnspod"] = LegoDnspod
	}

	if AlidnsEnabled {
		alidns_accesskey := viper.GetString("ALICLOUD_ACCESS_KEY")
		alidns_secretkey := viper.GetString("ALICLOUD_SECRET_KEY")
		alidns_email := viper.GetString("ADMIN_EMAIL")
		LegoAliyun := alidnsprovider.NewDefaultClient(
			alidns_email,
			alidns_accesskey,
			alidns_secretkey)

		Providers["alidns"] = LegoAliyun
	}

}

func init() {
	viperx.Default()
	// viperx.AddConfigPaths("$HOME/lego-certmgr")
}
