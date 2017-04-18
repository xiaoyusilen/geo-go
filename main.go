package main

import (
	"fmt"
	"net/http"
	"os"
	"runtime/debug"

	log "github.com/Sirupsen/logrus"
	"github.com/xiaoyusilen/geo-go/config"
	"github.com/xiaoyusilen/geo-go/route"
)

func initLogger(cfg *config.Config) {

	switch {
	case cfg.LogConfig.Formatter == config.LogTextFormatter:
		log.SetFormatter(&log.TextFormatter{})
	case cfg.LogConfig.Formatter == config.LogJSONFormatter:
		log.SetFormatter(&log.JSONFormatter{})
	}

	switch {
	case cfg.LogConfig.Output == config.LogConsoleOutput:
		log.SetOutput(os.Stdout)
	case cfg.LogConfig.Output == config.LogFileOutput:
		f, err := os.OpenFile(cfg.LogConfig.FilePath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
		if err != nil {
			log.Panicf("error opening log file: %v", err)
		}

		log.SetOutput(f)
	}

	switch {
	case cfg.LogConfig.Level == config.LogPanicLevel:
		log.SetLevel(log.PanicLevel)
	case cfg.LogConfig.Level == config.LogFatalLevel:
		log.SetLevel(log.FatalLevel)
	case cfg.LogConfig.Level == config.LogErrorLevel:
		log.SetLevel(log.ErrorLevel)
	case cfg.LogConfig.Level == config.LogWarnLevel:
		log.SetLevel(log.WarnLevel)
	case cfg.LogConfig.Level == config.LogInfoLevel:
		log.SetLevel(log.InfoLevel)
	case cfg.LogConfig.Level == config.LogDebugLevel:
		log.SetLevel(log.DebugLevel)
	}
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			debug.PrintStack()
		}
	}()

	// 读取配置文件
	cfg := config.ParseFromFlags()

	// 初始化log
	initLogger(cfg)

	// 启动Http 服务
	api := route.NewRestApi(cfg)

	address := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	log.Printf("Start successfully[%s] \n", address)

	if err := http.ListenAndServe(address, api.Api.MakeHandler()); err != nil {
		log.Fatal("api: Listen server failed, ", err)
	}
}
