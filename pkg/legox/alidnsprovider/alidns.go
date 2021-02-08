package alidnsprovider

import (
	"github.com/go-acme/lego/v4/challenge/dns01"
	"github.com/go-acme/lego/v4/providers/dns/alidns"
	"github.com/sirupsen/logrus"
	"github.com/tangx/srv-lego-certmgr/pkg/legox"
)

type Client struct {
	Email      string
	AccessKey  string
	SecretKey  string
	Nameserver string
	provider   *alidns.DNSProvider
	cli        *legox.LegoxClient
	nsopt      dns01.ChallengeOption
}

func NewDefaultClient(email string, accessKey string, secretKey string) *Client {
	c := &Client{
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

func (c *Client) signProvider() {
	p, err := NewProvider(c.AccessKey, c.SecretKey)
	if err != nil {
		logrus.Fatal(err)
	}

	c.provider = p
}

func (c *Client) signLegoxClient() {
	c.cli = legox.NewClient(c.Email)
}

func (c *Client) Init() {
	c.Default()
	c.signLegoxClient()
	c.signProvider()
}

func (c *Client) Default() {
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

func (c *Client) ApplyCertificate(domains ...string) (legox.Certificate, error) {
	return c.cli.ApplyCertificate(domains...)
}
