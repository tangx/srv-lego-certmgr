package main

import (
	"github.com/tangx/goutils/logx"
	"github.com/tangx/srv-lego-certmgr/cmd/certmgr/cmd"
)

func main() {
	cmd.Execute()
}

func init() {
	logx.InitLogrus()
}
