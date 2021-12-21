package utils

import (
	"time"

	"github.com/tangx/srv-lego-certmgr/pkg/legox"
)

func PushCert(domain string, cert legox.Certificate) {

	mapper.Store(domain, cert)
}

func GetCert(domain string) (legox.Certificate, bool) {
	return mapper.Get(domain)

}

func ListValidCerts() map[string]time.Time {

	return mapper.ListValidCerts()
}

func ListAll() map[string]legox.Certificate {
	return mapper.ListAll()
}
