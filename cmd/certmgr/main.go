package main

import (
	"github.com/tangx/goutils/logx"
	"github.com/tangx/srv-lego-certmgr/cmd/certmgr/global"
)

func main() {
	// cmd.Execute()

	global.App.Run(
		global.Server(),
	)

}

func init() {
	logx.InitLogrus()
}
