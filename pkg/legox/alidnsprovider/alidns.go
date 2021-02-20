package alidnsprovider

import (
	"github.com/go-acme/lego/v4/challenge/dns01"
	"github.com/go-acme/lego/v4/providers/dns/alidns"
	"github.com/sirupsen/logrus"
	"github.com/tangx/srv-lego-certmgr/pkg/legox"
)

type Config struct {
	Email      string                `env:"email"`
	AccessKey  string                `env:"access_key"`
	SecretKey  string                `env:"secret_key"`
	Nameserver string                `env:"nameserver"`
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
	ali.Default()
	ali.signLegoxClient()
	ali.signProvider()
}

func (ali *Config) Default() {
	if ali.Nameserver == "" {
		ali.nsopt = legox.DefaultNSOpts
	} else {
		ali.nsopt = legox.SetNSOpts(ali.Nameserver)
	}

	if ali.AccessKey == "" || ali.SecretKey == "" {
		logrus.Fatal("alidns AccessKey or SecretKey is required")
	}
	if ali.Email == "" {
		logrus.Fatal("alidns Email is required")
	}
}

func (ali *Config) ApplyCertificate(domains ...string) (legox.Certificate, error) {
	return ali.cli.ApplyCertificate(domains...)
}
