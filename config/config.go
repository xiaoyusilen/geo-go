// author by @xiaoyusilen

package config

import (
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/alecthomas/kingpin"
)

const (
	defaultHost = "127.0.0.1" // default binding host IP
	defaultPort = "9000"      // default binding port

	defaultTile38Address  = "127.0.0.1:9851"  // default tile38 server address
	defaultRethinkAddress = "127.0.0.1:28015" // default rethink driver address

	LogTextFormatter = "text"
	LogJSONFormatter = "json"

	LogConsoleOutput = "console"
	LogFileOutput    = "file"

	LogPanicLevel = "panic"
	LogFatalLevel = "fatal"
	LogErrorLevel = "error"
	LogWarnLevel  = "warn"
	LogInfoLevel  = "info"
	LogDebugLevel = "debug"
)

// LogConfig is logging configuration
type LogConfig struct {
	Formatter string // log formatter. TextFormatter or JsonFormatter
	Output    string // log output, console or file
	FilePath  string // log file path, enabled when output is file
	Level     string // log level, debug/info/warn/error/fatal/panic
}

// Config is application startup configuration
type Config struct {
	Host string
	Port int

	LogConfig *LogConfig // 日志配置

	Tile38Address           string
	Tile38DBNumber          int
	Tile38MaxIdleConnection int
	Tile38IdleTimeout       time.Duration
	Tile38ConnectTimeout    time.Duration
	Tile38ReadTimeout       time.Duration
	Tile38WriteTimeout      time.Duration
	RethinkAddress          string
}

// SetFormatter sets the formatter of logger
func (logConfig *LogConfig) SetFormatter(formatter *string) {
	switch *formatter {
	case LogTextFormatter:
		logConfig.Formatter = LogTextFormatter
	case LogJSONFormatter:
		logConfig.Formatter = LogJSONFormatter
	default:
		logConfig.Formatter = LogTextFormatter
	}
}

// SetOutput sets the output of logger
func (logConfig *LogConfig) SetOutput(output *string, filePath *string) {

	switch *output {
	case LogConsoleOutput:
		logConfig.Output = LogConsoleOutput
	case LogFileOutput:
		logConfig.Output = LogFileOutput
	default:
		logConfig.Output = LogConsoleOutput
	}

	if *output == LogFileOutput {
		if *filePath == "" {
			log.Fatal("Log file path should be set when output is file")
		}

		logConfig.FilePath = *filePath
	}
}

// SetLevel sets the level of logger
func (logConfig *LogConfig) SetLevel(level *string) {
	switch *level {
	case LogPanicLevel:
		logConfig.Level = LogPanicLevel
	case LogFatalLevel:
		logConfig.Level = LogFatalLevel
	case LogErrorLevel:
		logConfig.Level = LogErrorLevel
	case LogWarnLevel:
		logConfig.Level = LogWarnLevel
	case LogInfoLevel:
		logConfig.Level = LogInfoLevel
	case LogDebugLevel:
		logConfig.Level = LogDebugLevel
	default:
		logConfig.Level = LogErrorLevel
	}
}

func ParseFromFlags() *Config {

	// Parse flags
	host := kingpin.Flag("host", "Set service listening host address.").Default(defaultHost).String()
	port := kingpin.Flag("port", "Set service listening port.").Default(defaultPort).Int()

	logFormatter := kingpin.Flag("logFormatter", "Log formatter").Default(LogTextFormatter).String()
	logOutput := kingpin.Flag("logOutput", "Log output").Default(LogConsoleOutput).String()
	logFilePath := kingpin.Flag("logFilePath", "Log file path").String()
	logLevel := kingpin.Flag("logLevel", "Log level").Default(LogDebugLevel).String()

	tile38address := kingpin.Flag("tile38Address", "Tile38 Server address").Default(defaultTile38Address).String()
	tile38DBNumber := kingpin.Flag("tile38DBNumber", "Tile38 DB number.").Default("1").Int()
	tile38MaxIdleConnection := kingpin.Flag("tile38MaxIdleConnection", "Tile38 max idle connection.").Default("40").Int()
	tile38IdleTimeout := kingpin.Flag("tile38IdleTimeout", "Tile38 idle timeout.").Default("180").Int64()
	tile38ConnectTimeout := kingpin.Flag("tile38ConnectTimeout", "Tile38 connect timeout.").Default("10").Int64()
	tile38ReadTimeout := kingpin.Flag("tile38ReadTimeout", "Tile38 read timeout.").Default("10").Int64()
	tile38WriteTimeout := kingpin.Flag("tile38WriteTimeout", "Tile38 write timeout.").Default("10").Int64()

	rethinkaddress := kingpin.Flag("rethinkAddress", "Rethink Driver address").Default(defaultRethinkAddress).String()

	kingpin.Parse()

	// initialize startup configuration
	cfg := Config{
		Host:                    *host,
		Port:                    *port,
		Tile38Address:           *tile38address,
		Tile38DBNumber:          *tile38DBNumber,
		Tile38MaxIdleConnection: *tile38MaxIdleConnection,
		Tile38IdleTimeout:       time.Duration(*tile38IdleTimeout) * time.Second,
		Tile38ConnectTimeout:    time.Duration(*tile38ConnectTimeout) * time.Second,
		Tile38ReadTimeout:       time.Duration(*tile38ReadTimeout) * time.Second,
		Tile38WriteTimeout:      time.Duration(*tile38WriteTimeout) * time.Second,

		RethinkAddress: *rethinkaddress,

		LogConfig: &LogConfig{},
	}

	cfg.LogConfig.SetFormatter(logFormatter)
	cfg.LogConfig.SetOutput(logOutput, logFilePath)
	cfg.LogConfig.SetLevel(logLevel)

	return &cfg
}
