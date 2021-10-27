package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/Chipazawra/czwr-mailing-auth/pkg/pprofwrapper"
	_ "github.com/Chipazawra/czwr-mailing-profile/doc"
	mongoctx "github.com/Chipazawra/czwr-mailing-profile/internal/dbcontext/mongo"
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

	var host, port, dbuser, dbpass, dbclst string

	flag.StringVar(&host, "host", "", "Host on which to start listening")
	flag.StringVar(&port, "port", "", "Port on which to start listening")
	flag.StringVar(&dbuser, "dbuser", "", "db user")
	flag.StringVar(&dbpass, "dbpass", "", "db pass")
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

	if dbuser == "" {
		dbuser = os.Getenv("DB_USER")
		if dbuser == "" {
			panic("db user is not set, use env=\"DB_USER\" or cmd args \"-dbuser\".")
		}
	}

	if dbpass == "" {
		dbpass = os.Getenv("DB_PASS")
		if dbpass == "" {
			panic("db pass is not set, use env=\"DB_PASS\" or cmd args \"-dbpass\".")
		}
	}

	if dbclst == "" {
		dbclst = os.Getenv("DB_CLST")
		if dbclst == "" {
			panic("db cluser is not set, use env=\"DB_CLST\" or cmd args \"-dbclst\".")
		}
	}

	ctx := context.TODO()
	//init mongo
	mClient := mongoctx.New()
	err := mClient.Connect(ctx, dbuser, dbpass, dbclst)
	defer mClient.Disonnect(ctx)
	if err != nil {
		panic(err)
	}

	httpEngine := gin.New()
	httpEngine.Use(gin.Recovery())
	// log template
	httpEngine.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		Formatter: logfmt,
		Output:    os.Stdout,
	}))

	service := profile.New(mClient)
	group := service.Register(httpEngine)
	//profiler
	pprofrp := pprofwrapper.New()
	pprofrp.Register(group)
	//doc
	group.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err = httpEngine.Run(fmt.Sprintf("%v:%v", host, port))

	if err != nil {
		panic(err)
	}
}

func logfmt(params gin.LogFormatterParams) string {

	var statusColor, methodColor, resetColor string
	if params.IsOutputColor() {
		statusColor = params.StatusCodeColor()
		methodColor = params.MethodColor()
		resetColor = params.ResetColor()
	}

	if params.Latency > time.Minute {
		// Truncate in a golang < 1.8 safe way
		params.Latency = params.Latency - params.Latency%time.Second
	}

	return fmt.Sprintf("[PROFILE-LOG] %v |%s %3d %s| %13v | %15s |%s %-7s %s %#v\n%s",
		params.TimeStamp.Format("2006/01/02 - 15:04:05"),
		statusColor, params.StatusCode, resetColor,
		params.Latency,
		params.ClientIP,
		methodColor, params.Method, resetColor,
		params.Path,
		params.ErrorMessage,
	)
}
