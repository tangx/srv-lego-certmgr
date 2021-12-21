package global

import (
	"github.com/go-jarvis/jarvis"
	"github.com/go-jarvis/jarvis/pkg/appctx"
	"github.com/go-jarvis/rum-gonic/confhttp"
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
