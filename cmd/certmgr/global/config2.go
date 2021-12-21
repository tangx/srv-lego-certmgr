package global

import (
	"github.com/go-jarvis/jarvis"
	"github.com/go-jarvis/jarvis/pkg/appctx"
	"github.com/go-jarvis/rum-gonic/confhttp"
	"github.com/tangx/srv-lego-certmgr/cmd/certmgr/apis"
	"github.com/tangx/srv-lego-certmgr/pkg/legox/alidnsprovider"
	"github.com/tangx/srv-lego-certmgr/pkg/legox/dnspodprovider"
)

var (
	server = &confhttp.Server{}
	alidns = &alidnsprovider.Config{}
	dnspod = &dnspodprovider.Config{}
)

var (
	App = jarvis.New().WithOptions(
		appctx.WithName(Appname),
		appctx.WithRoot("../.."),
	)
)

func init() {
	config := &struct {
		HttpServer *confhttp.Server
		Alidns     *alidnsprovider.Config
		Dnspod     *dnspodprovider.Config
	}{
		HttpServer: server,
		Alidns:     alidns,
		Dnspod:     dnspod,
	}

	_ = App.Conf(config)
}

func Server() *confhttp.Server {
	server.Register(apis.BaseRouterGroup)
	return server
}
