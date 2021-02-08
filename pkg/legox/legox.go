package legox

import (
	"fmt"

	"github.com/go-acme/lego/v4/certcrypto"
	"github.com/go-acme/lego/v4/certificate"
	"github.com/go-acme/lego/v4/challenge"
	"github.com/go-acme/lego/v4/challenge/dns01"
	"github.com/go-acme/lego/v4/lego"
	"github.com/go-acme/lego/v4/registration"
	"github.com/sirupsen/logrus"
)

type LegoxClient struct {
	user   *user
	client *lego.Client
}

func NewClient(email string) *LegoxClient {

	u := user{Email: email}
	u.initial()

	lx := &LegoxClient{
		user: &u,
	}

	lx.initial()
	return lx
}

func (lx *LegoxClient) initial() {

	// 生成新客户端
	config := lego.NewConfig(lx.user)
	config.Certificate.KeyType = certcrypto.RSA2048

	cli, err := lego.NewClient(config)
	if err != nil {
		logrus.Fatal(err)
	}

	lx.client = cli

	// 注册新用户
	if lx.user.Registration == nil {
		reg, err := lx.client.Registration.Register(
			registration.RegisterOptions{
				TermsOfServiceAgreed: true,
			})

		if err != nil {
			logrus.Fatal(err)
		}

		lx.user.register(reg)
	}
}

func (lx *LegoxClient) ApplyCertificate(domains ...string) (Certificate, error) {
	request := certificate.ObtainRequest{
		Domains: domains,
		Bundle:  true,
	}

	certs, err := lx.client.Certificate.Obtain(request)
	if err != nil {
		// logrus.Fatal(err)
		logrus.Error(err)

		return Certificate{}, err
	}

	c := Certificate{
		Domain:        certs.Domain,
		CertURL:       certs.CertURL,
		CertStableURL: certs.CertStableURL,
		PrivateKey:    fmt.Sprintf("%s", certs.PrivateKey),
		Certificate:   fmt.Sprintf("%s", certs.Certificate),
	}

	c.SetValidationTime()

	return c, nil
}

func (lx *LegoxClient) SetDNS01Provider(p challenge.Provider, opts ...dns01.ChallengeOption) {
	err := lx.client.Challenge.SetDNS01Provider(p, opts...)
	if err != nil {
		logrus.Fatal(err)
	}
}
