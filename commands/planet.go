// Copyright 2017 Sergiu Bodiu
//
// Use of this source code is governed by and Apache2
// license that can be found in the LICENSE file
package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var RootCmd = &cobra.Command{
	Use:   "dagobah",
	Short: "Dagobah is an awesome planet style RSS aggregator",
	Long: `Dagobah provides planet style RSS aggregation. It
is inspired by python planet. It has a simple YAML configuration
and provides it's own webserver.`,
	Run: rootRun,
}

var CfgFile string

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVar(&CfgFile, "config", "", "config file (default is $HOME/dagobah/config.yaml)")
}

func initConfig() {
	if CfgFile != "" {
		viper.SetConfigFile(CfgFile)
	}
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/dagobah")
	viper.AddConfigPath("$HOME/.dagobah/")
	viper.ReadInConfig()
}

func rootRun(cmd *cobra.Command, args []string) {
	fmt.Println(viper.Get("feeds"))
	fmt.Println(viper.GetString("appname"))
}

func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}