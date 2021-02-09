package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tangx/srv-lego-certmgr/cmd/certmgr/global"
	"github.com/tangx/srv-lego-certmgr/cmd/certmgr/routes"
)

var rootCmd = &cobra.Command{
	Use: "certmgr",
	Run: func(cmd *cobra.Command, args []string) {

		global.Initial()

		routes.AppendRoute(global.Server)
		_ = global.Server.Run(":80")

	},
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&global.AlidnsEnabled, "alidns", "", false, "enabled/disabled alidns provider")
	rootCmd.PersistentFlags().BoolVarP(&global.DnspodEnabled, "dnspod", "", false, "enabled/disabled dnspod provider")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
