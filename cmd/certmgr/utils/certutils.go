package utils

import (
	"github.com/tangx/srv-lego-certmgr/cmd/certmgr/global"
	"github.com/tangx/srv-lego-certmgr/pkg/legox"
)

func PushCert(domain string, cert legox.Certificate) {
	global.CertMap[domain] = cert
}

func GetCert(domain string) (legox.Certificate, bool) {
	cert, ok := global.CertMap[domain]
	if ok {
		return cert, true
	}

	return legox.Certificate{}, false
}
