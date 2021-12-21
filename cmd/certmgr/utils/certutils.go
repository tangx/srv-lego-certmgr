package utils

import (
	"time"

	"github.com/tangx/srv-lego-certmgr/pkg/legox"
)

// 用于保存生成的证书，方便 GET 时快速返回。 不持久化
var certMap = make(map[string](legox.Certificate))

func PushCert(domain string, cert legox.Certificate) {
	certMap[domain] = cert
}

func GetCert(domain string) (legox.Certificate, bool) {
	cert, ok := certMap[domain]
	if ok {
		return cert, true
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
