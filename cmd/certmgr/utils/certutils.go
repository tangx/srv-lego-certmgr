package utils

import (
	"time"

	"github.com/spf13/viper"
	"github.com/tangx/srv-lego-certmgr/pkg/legox"
	"github.com/tangx/srv-lego-certmgr/pkg/storage"
	"github.com/tangx/srv-lego-certmgr/pkg/storage/filesystem"
)

// 用于保存生成的证书，方便 GET 时快速返回。
var (
	certMap  = make(map[string](legox.Certificate))
	storager storage.Storager
)

// initial backend storage
func init() {
	// BackendStorageClass
	class := viper.GetString("BackendStorageClass")
	switch class {
	case "filesystem":
		//todo
		dir := viper.GetString("StorageFilesystemDir")
		storager = filesystem.NewStorager(dir)
	}
}

func PushCert(domain string, cert legox.Certificate) {
	certMap[domain] = cert

	if storager != nil {
		storager.Store(&cert)
	}
}

func GetCert(domain string) (legox.Certificate, bool) {
	cert, ok := certMap[domain]
	if ok {
		return cert, true
	}

	if storager != nil {
		crt := storager.GetByName(domain)
		if crt != nil {
			certMap[domain] = *crt

			return *crt, true
		}
	}

	return legox.Certificate{}, false
}

func ListValidCerts() map[string]time.Time {

	m := make(map[string]time.Time)
	for _, cert := range certMap {
		m[cert.Domain] = cert.NotAfter
	}

	return m
}
func ListAll() map[string]legox.Certificate {
	return certMap
}
