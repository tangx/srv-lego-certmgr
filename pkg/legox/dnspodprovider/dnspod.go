package dnspodprovider

import (
	"github.com/go-acme/lego/v4/challenge/dns01"
	"github.com/go-acme/lego/v4/providers/dns/dnspod"
	"github.com/sirupsen/logrus"
	"github.com/tangx/srv-lego-certmgr/pkg/legox"
)

// NewProvider 返回一个 dnspod provider
func NewProvider(token string) *dnspod.DNSProvider {
	config := dnspod.NewDefaultConfig()
	config.LoginToken = token

	p, err := dnspod.NewDNSProviderConfig(config)
	if err != nil {
		logrus.Fatal(err)
	}
	return p
}

type Config struct {
	nickName   string
	Enabled    bool                  `env:""`
	Email      string                `env:""`
	Token      string                `env:""`
	Nameserver string                `env:""`
	provider   *dnspod.DNSProvider   `env:"-"`
	cli        *legox.LegoxClient    `env:"-"`
	nsopt      dns01.ChallengeOption `env:"-"`
}

func NewDefualtClient(email string, token string) *Config {
	dp := &Config{
		Email: email,
		Token: token,
	}

	dp.Init()

	return dp
}

// 初始化
func (dp *Config) Init() {

	dp.SetDefaults()

	if dp.Enabled {
		dp.signLegoxClient()
		dp.signProvider()
	}
}

// 初始化默认信息
func (dp *Config) SetDefaults() {

	if dp.nickName == "" {
		dp.nickName = "dnspod"
	}

	// 设置 Nameserver 信息
	if dp.Nameserver == "" {
		dp.nsopt = legox.DefaultNSOpts
	} else {
		dp.nsopt = legox.SetNSOpts(dp.Nameserver)
	}
}

func (dp *Config) NickName() string {
	return dp.nickName
}

// ApplyCertificate 向 letsencrypt 申请证书
func (dp *Config) ApplyCertificate(domains ...string) (legox.Certificate, error) {
	return dp.cli.ApplyCertificate(domains...)
}

// signLegoxClient 将初始化一个 legox 客户端
func (dp *Config) signLegoxClient() {
	dp.cli = legox.NewClient(dp.Email)
}

// signProvider 向 legox 中加入 provider 的信息
func (dp *Config) signProvider() {
	dp.provider = NewProvider(dp.Token)
	dp.cli.SetDNS01Provider(dp.provider, dp.nsopt)
}
