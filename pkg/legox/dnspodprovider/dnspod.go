package dnspodprovider

import (
	"github.com/go-acme/lego/v4/challenge/dns01"
	"github.com/go-acme/lego/v4/providers/dns/dnspod"
	"github.com/sirupsen/logrus"
	"github.com/tangx/srv-lego-certmgr/pkg/legox"
)

type Dnspod struct {
	Email      string                `env:"email"`
	Token      string                `env:"token"`
	Nameserver string                `env:"nameserver"`
	provider   *dnspod.DNSProvider   `env:"-"`
	cli        *legox.LegoxClient    `env:"-"`
	nsopt      dns01.ChallengeOption `env:"-"`
}

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

func NewDefualtClient(email string, token string) *Dnspod {
	dp := &Dnspod{
		Email: email,
		Token: token,
	}

	dp.Init()

	return dp
}

// 初始化
func (dp *Dnspod) Init() {
	dp.Default()
	dp.signLegoxClient()
	dp.signProvider()
}

// 初始化默认信息
func (dp *Dnspod) Default() {

	// 设置 Nameserver 信息
	if dp.Nameserver == "" {
		dp.nsopt = legox.DefaultNSOpts
	} else {
		dp.nsopt = legox.SetNSOpts(dp.Nameserver)
	}

	if dp.Token == "" {
		logrus.Fatal("dnspod token is missing")
	}

	if dp.Email == "" {
		logrus.Fatal("user email is required")
	}

}

// ApplyCertificate 向 letsencrypt 申请证书
func (dp *Dnspod) ApplyCertificate(domains ...string) (legox.Certificate, error) {
	return dp.cli.ApplyCertificate(domains...)
}

// signLegoxClient 将初始化一个 legox 客户端
func (dp *Dnspod) signLegoxClient() {
	dp.cli = legox.NewClient(dp.Email)
}

// signProvider 向 legox 中加入 provider 的信息
func (dp *Dnspod) signProvider() {
	dp.provider = NewProvider(dp.Token)
	dp.cli.SetDNS01Provider(dp.provider, dp.nsopt)
}
