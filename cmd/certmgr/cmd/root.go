package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tangx/srv-lego-certmgr/cmd/certmgr/global"
	"github.com/tangx/srv-lego-certmgr/cmd/certmgr/routes"
	"github.com/tangx/srv-lego-certmgr/version"
)

var rootCmd = &cobra.Command{
	Use:  "certmgr",
	Long: fmt.Sprintf("Version:\n  v%s", version.Version),
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
	Run: func(cmd *cobra.Command, args []string) {
		global.Initial()
		routes.AppendRoute(global.Server)
		_ = global.Server.Run(listenAddr)
	},
}

var listenAddr string

func init() {
	rootCmd.PersistentFlags().BoolVarP(&global.AlidnsEnabled, "alidns", "", false, "enabled/disabled alidns provider")
	rootCmd.PersistentFlags().BoolVarP(&global.DnspodEnabled, "dnspod", "", false, "enabled/disabled dnspod provider")
	rootCmd.PersistentFlags().StringVarP(&listenAddr, "listen", "", ":8080", "服务运行监听地址端口")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
