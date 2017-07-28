// Copyright 2017 Sergiu Bodiu
//
// Use of this source code is governed by and Apache2
// license that can be found in the LICENSE file
package commands

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

    "github.com/GeertJohan/go.rice"
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
	templates := loadTemplates("full.html")
	r.SetHTMLTemplate(templates)

	r.GET("/ping", func(context *gin.Context) {
		context.String(200, "pong")
	})

	r.GET("/", homeRoute)
	r.GET("/static/*filepath", staticServe)

	port := viper.GetString("port")
	fmt.Println("Running on port:", port)
	r.Run(":" + port)
}

func homeRoute(c *gin.Context) {
	obj := gin.H{"title": "Go Rules"}
	c.HTML(200, "full.html", obj)
}

func staticServe(c *gin.Context) {
	static, err := rice.FindBox("static")
	if err != nil {
		log.Fatal(err)
	}

	original := c.Request.URL.Path
	c.Request.URL.Path = c.Params.ByName("filepath")
	fmt.Println(c.Params.ByName("filepath"))
	http.FileServer(static.HTTPBox()).ServeHTTP(c.Writer, c.Request)
	c.Request.URL.Path = original
}

func loadTemplates(list ...string) *template.Template {
	templateBox, err := rice.FindBox("templates")
	if err != nil {
		log.Fatal(err)
	}
	templates := template.New("")

	for _, x := range list {
		templateString, err := templateBox.String(x)
		if err != nil {
			log.Fatal(err)
		}

		// get file contents as string
		_, err = templates.New(x).Parse(templateString)
		if err != nil {
			log.Fatal(err)
		}
	}

	return templates
}