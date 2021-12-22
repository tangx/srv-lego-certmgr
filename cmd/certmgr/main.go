package main

import (
	"github.com/tangx/goutils/logx"
	"github.com/tangx/srv-lego-certmgr/cmd/certmgr/apis"
	"github.com/tangx/srv-lego-certmgr/cmd/certmgr/global"
)

func main() {

	// global.Server().StaticFile("/index", "./static/index.html")
	// global.Server().Register(apis.BaseRouterGroup)
	global.Server().Register(apis.RootRouterGroup)
	global.App.Run(
		global.Server(),
	)

}

func init() {
	logx.InitLogrus()
}
