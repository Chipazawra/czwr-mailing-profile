package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Chipazawra/czwr-mailing-auth/pkg/pprofwrapper"
	_ "github.com/Chipazawra/czwr-mailing-profile/doc"
	"github.com/Chipazawra/czwr-mailing-profile/internal/profile"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

var (
	defaultHost = "0.0.0.0"
	defaultPort = "8884"
)

// @title czwrMailing - profile service
// @version 1.0
// @description This is a sample mailing servivce.
func main() {

	var host, port string

	flag.StringVar(&host, "host", "", "Host on which to start listening")
	flag.StringVar(&port, "port", "", "Port on which to start listening")
	flag.Parse()

	if host == "" {
		host = os.Getenv("AUTH_HOST")
		if host == "" {
			host = defaultHost
		}
	}

	if port == "" {
		port = os.Getenv("AUTH_PORT")
		if port == "" {
			port = defaultPort
		}
	}

	httpEngine := gin.New()
	httpEngine.Use(gin.Recovery())

	service := profile.New()
	group := service.Register(httpEngine)
	//profiler
	pprofrp := pprofwrapper.New()
	pprofrp.Register(group)
	//doc
	group.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := httpEngine.Run(fmt.Sprintf("%v:%v", host, port))

	if err != nil {
		panic(err)
	}

}
