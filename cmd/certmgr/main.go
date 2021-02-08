package main

import (
	"github.com/tangx/srv-lego-certmgr/cmd/certmgr/global"
	"github.com/tangx/srv-lego-certmgr/cmd/certmgr/routes"
)

func main() {
	global.Server.Run(":80")
}

func init() {
	routes.AppendRoute(global.Server)
}
