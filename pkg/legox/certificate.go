package legox

import (
	"time"

	"github.com/sirupsen/logrus"
	"github.com/tangx/srv-lego-certmgr/pkg/x509util"
)

type Certificate struct {
	Domain string `json:"domain,omitempty"`
	// acem 下载证书地址
	CertURL       string    `json:"cert_url,omitempty"`
	CertStableURL string    `json:"cert_stable_url,omitempty"`
	PrivateKey    string    `json:"private_key,omitempty"`
	Certificate   string    `json:"certificate,omitempty"`
	NotAfter      time.Time `json:"not_after,omitempty"`
	NotBefore     time.Time `json:"not_before,omitempty"`
}

// SetValidationTime 填充证书有效时间
func (c *Certificate) SetValidationTime() {
	if c.Certificate == "" {
		return
	}

	x509cert, err := x509util.ParseCertificate(c.Certificate)
	if err != nil {
		logrus.Error(err)

		return
	}

	c.NotAfter = x509cert.NotAfter.Local()
	c.NotBefore = x509cert.NotBefore.Local()

}
