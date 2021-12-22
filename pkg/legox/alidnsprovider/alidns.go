package alidnsprovider

import (
	"github.com/go-acme/lego/v4/challenge/dns01"
	"github.com/go-acme/lego/v4/providers/dns/alidns"
	"github.com/sirupsen/logrus"
	"github.com/tangx/srv-lego-certmgr/pkg/legox"
)

type Config struct {
	nickName   string
	Enabled    bool                  `env:""`
	Email      string                `env:""`
	AccessKey  string                `env:""`
	SecretKey  string                `env:""`
	Nameserver string                `env:""`
	provider   *alidns.DNSProvider   `env:"-"`
	cli        *legox.LegoxClient    `env:"-"`
	nsopt      dns01.ChallengeOption `env:"-"`
}

func NewDefaultClient(email string, accessKey string, secretKey string) *Config {
	ali := &Config{
		Email:     email,
		AccessKey: accessKey,
		SecretKey: secretKey,
	}
	ali.Init()

	return ali
}

func NewProvider(access, secret string) (*alidns.DNSProvider, error) {
	config := alidns.NewDefaultConfig()
	config.APIKey = access
	config.SecretKey = secret

	return alidns.NewDNSProviderConfig(config)
}

func (ali *Config) NickName() string {
	return ali.nickName
}

func (ali *Config) signProvider() {
	p, err := NewProvider(ali.AccessKey, ali.SecretKey)
	if err != nil {
		logrus.Fatal(err)
	}

	ali.provider = p
	ali.cli.SetDNS01Provider(ali.provider, ali.nsopt)
}

func (ali *Config) signLegoxClient() {
	ali.cli = legox.NewClient(ali.Email)
}

func (ali *Config) Init() {
	ali.SetDefaults()

	if ali.Enabled {
		ali.signLegoxClient()
		ali.signProvider()
	}
}

func (ali *Config) SetDefaults() {
	if ali.nickName == "" {
		ali.nickName = "alidns"
	}

	if ali.Nameserver == "" {
		ali.nsopt = legox.DefaultNSOpts
	} else {
		ali.nsopt = legox.SetNSOpts(ali.Nameserver)
	}

}

func (ali *Config) ApplyCertificate(domains ...string) (legox.Certificate, error) {
	return ali.cli.ApplyCertificate(domains...)
}
