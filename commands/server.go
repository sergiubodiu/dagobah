// Copyright 2017 Sergiu Bodiu
//
// Use of this source code is governed by and Apache2
// license that can be found in the LICENSE file
package commands

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/gin-gonic/gin"
)

var serverCmd = &cobra.Command{
	Use: "server",
	Short: "Server for feeds",
	Long: `Dagobah will serve all feeds listed in the config file.`,
	Run: serverRun,
}

func init()  {
	serverCmd.Flags().Int("port", 1138, "Port to run Dagobah server on")
	viper.BindPFlag("port", serverCmd.Flags().Lookup("port"))
}

func serverRun(cmd *cobra.Command, args []string) {
	r := gin.Default()

	r.GET("/ping", func(context *gin.Context) {
		context.String(200, "pong")
	})

	port := viper.GetString("port")
	fmt.Println("Running on port:", port)
	r.Run(":" + port)
}