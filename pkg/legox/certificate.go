package legox

import (
	"crypto/x509"
	"encoding/pem"
	"time"

	"github.com/sirupsen/logrus"
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

	data := []byte(c.Certificate)

	// https://blog.csdn.net/u011228889/article/details/81480617
	blk, _ := pem.Decode(data)
	if blk == nil {
		logrus.Error("pem decode failed: ")
		return
	}

	x509cert, err := x509.ParseCertificate(blk.Bytes)
	if err != nil {
		logrus.Error(err)
		return
	}

	c.NotAfter = x509cert.NotAfter.Local()
	c.NotBefore = x509cert.NotBefore.Local()

}
