package alidnsprovider

import (
	"github.com/go-acme/lego/v4/challenge/dns01"
	"github.com/go-acme/lego/v4/providers/dns/alidns"
	"github.com/sirupsen/logrus"
	"github.com/tangx/srv-lego-certmgr/pkg/legox"
)

type Alidns struct {
	Email      string                `env:"email"`
	AccessKey  string                `env:"access_key"`
	SecretKey  string                `env:"secret_key"`
	Nameserver string                `env:"nameserver"`
	provider   *alidns.DNSProvider   `env:"-"`
	cli        *legox.LegoxClient    `env:"-"`
	nsopt      dns01.ChallengeOption `env:"-"`
}

func NewDefaultClient(email string, accessKey string, secretKey string) *Alidns {
	c := &Alidns{
		Email:     email,
		AccessKey: accessKey,
		SecretKey: secretKey,
	}
	c.Init()

	return c
}

func NewProvider(access, secret string) (*alidns.DNSProvider, error) {
	config := alidns.NewDefaultConfig()
	config.APIKey = access
	config.SecretKey = secret

	return alidns.NewDNSProviderConfig(config)
}

func (c *Alidns) signProvider() {
	p, err := NewProvider(c.AccessKey, c.SecretKey)
	if err != nil {
		logrus.Fatal(err)
	}

	c.provider = p
	c.cli.SetDNS01Provider(c.provider, c.nsopt)
}

func (c *Alidns) signLegoxClient() {
	c.cli = legox.NewClient(c.Email)
}

func (c *Alidns) Init() {
	c.Default()
	c.signLegoxClient()
	c.signProvider()
}

func (c *Alidns) Default() {
	if c.Nameserver == "" {
		c.nsopt = legox.DefaultNSOpts
	} else {
		c.nsopt = legox.SetNSOpts(c.Nameserver)
	}

	if c.AccessKey == "" || c.SecretKey == "" {
		logrus.Fatal("alidns AccessKey or SecretKey is required")
	}
	if c.Email == "" {
		logrus.Fatal("alidns Email is required")
	}
}

func (c *Alidns) ApplyCertificate(domains ...string) (legox.Certificate, error) {
	return c.cli.ApplyCertificate(domains...)
}
