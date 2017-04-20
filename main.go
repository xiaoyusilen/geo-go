package main

import (
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
	// 读取配置文件
	cfg := config.ParseFromFlags()

	// 初始化log
	initLogger(cfg)

	// 初始化Http服务
	r := route.HandleRest(cfg)

	r.Router.Run(":9000")

	defer func() {
		if err := recover(); err != nil {
			debug.PrintStack()
		}

		// 关闭tile38连接
		r.Tile38.Close()

		// 关闭Rethinkdb连接
		r.Rethink.Close()

		// TODO: 关闭Http service
	}()

}
