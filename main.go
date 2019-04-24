package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"time"

	"github.com/eaglerayp/go-api-template/apiserver"
	"github.com/spf13/viper"
)

type options struct {
	version    bool
	prof       bool
	configFile string
}

var opts options
var systemCtx context.Context

func init() {
	flag.BoolVar(&opts.version, "build", false, "GoLang build version.")
	flag.BoolVar(&opts.prof, "prof", false, "GoLang profiling function.")
	flag.StringVar(&opts.configFile, "config", "./local.toml", "Path to Config File")

	systemCtx = context.Background()
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [arguments] <command> \n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()
	if opts.version {
		fmt.Fprintf(os.Stderr, "%s\n", runtime.Version())
	}
	if opts.prof {
		ActivateProfile()
	}
	viper.AutomaticEnv()
	viper.SetConfigFile(opts.configFile)
	viper.ReadInConfig()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	log.SetFlags(log.LstdFlags)

	httpServer := apiserver.InitGinServer(systemCtx)

	// graceful shutdown
	quit := make(chan os.Signal, 5)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
