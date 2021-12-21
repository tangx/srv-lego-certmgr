package apis

import (
	"github.com/go-jarvis/rum-gonic/rum"
	"github.com/tangx/srv-lego-certmgr/cmd/certmgr/apis/query"
)

var BaseRouterGroup = rum.NewRouterGroup("/lego-certmgr")

func init() {
	// todo: register sub routes

	BaseRouterGroup.Register(query.QueryRouterGroup)
}
