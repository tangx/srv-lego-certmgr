package global

import (
	"github.com/go-jarvis/jarvis"
	"github.com/go-jarvis/jarvis/pkg/appctx"
	"github.com/go-jarvis/rum-gonic/confhttp"
	"github.com/tangx/srv-lego-certmgr/cmd/certmgr/apis"
)

var (
	server = &confhttp.Server{}

	App = jarvis.New().WithOptions(
		appctx.WithName(Appname),
		appctx.WithRoot("../.."),
	)
)

func init() {
	config := &struct {
		HttpServer *confhttp.Server
	}{
		HttpServer: server,
	}

	_ = App.Conf(config)
}

func Server() *confhttp.Server {
	server.Register(apis.BaseRouterGroup)
	return server
}
