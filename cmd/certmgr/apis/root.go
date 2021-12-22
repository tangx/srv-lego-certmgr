package apis

import (
	"github.com/go-jarvis/rum-gonic/pkg/middlewares"
	"github.com/go-jarvis/rum-gonic/rum"
	"github.com/tangx/srv-lego-certmgr/cmd/certmgr/apis/gen"
	"github.com/tangx/srv-lego-certmgr/cmd/certmgr/apis/query"
)

var RootRouterGroup = rum.NewRouterGroup("/")
var BaseRouterGroup = rum.NewRouterGroup("/lego-certmgr")

func init() {

	{
		BaseRouterGroup.Register(middlewares.DefaultNoCacheIndex())
		RootRouterGroup.StaticFile("/", "./static/index.html")
		RootRouterGroup.Register(BaseRouterGroup)
	}

	// todo: register sub routes
	BaseRouterGroup.Register(query.QueryRouterGroup)
	BaseRouterGroup.Register(gen.GenRouterGroup)
}
