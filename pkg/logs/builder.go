package logs

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/seerx/base/pkg/logs/transfers"
	"github.com/sirupsen/logrus"
)

// Builder 日志 Builder
type Builder struct {
	appTag          string       // 应用标志
	prettyJson      bool         // 输出格式化的 json
	timestampFormat string       // 日期格式
	reportCaller    bool         // 是否输出日志发生地址 , 文件 函数 行号
	level           logrus.Level // 日志输出级别
	outputJson      bool

	console bool   // 在控制台输出
	udpHost string // 接收日志的 udp 主机
	udpPort int    // 接收日志的 udp 端口
}

// NewBuilder 创建 builder
func NewBuilder() *Builder {
	return &Builder{
		console:         true,
		level:           logrus.InfoLevel,
		timestampFormat: time.RFC3339,
		reportCaller:    true,
	}
}

func (b *Builder) Build() *logrus.Logger {
	var logger = logrus.New()
	setNull(logger)
	if b.outputJson {
		logger.Formatter = &logrus.JSONFormatter{
			TimestampFormat:  b.timestampFormat,
			DisableTimestamp: false,
			DataKey:          "",
			FieldMap:         nil,
			CallerPrettyfier: nil,
			PrettyPrint:      b.prettyJson,
		}
	} else {
		logger.Formatter = &TextFormatter{
			timeFormat: b.timestampFormat,
		}
	}
	logger.Level = b.level
	logger.ReportCaller = b.reportCaller

	var txfns []transfers.TransferFn
	if b.console {
		txfns = append(txfns, MakeTransfer(nil))
	}
	if b.udpHost != "" && b.udpPort > 0 {
		txfns = append(txfns, MakeTransfer(&transfers.TransferConfigure{
			Type:   transfers.UDP,
			Server: b.udpHost,
			Port:   b.udpPort,
		}))
	}

	logger.AddHook(NewTransferHook(b.appTag,
		logger.Formatter,
		txfns...))

	return logger
}

// ReportCaller 是否报告日志地址
func (b *Builder) ReportCaller(report bool) *Builder {
	b.reportCaller = report
	return b
}

// Level 日志级别
func (b *Builder) Level(level logrus.Level) *Builder {
	b.level = level
	return b
}

// WriteToUDP 设置日志输出到 udp
func (b *Builder) WriteToUDP(host string, port int) *Builder {
	b.udpHost = host
	b.udpPort = port
	return b
}

// WriteToConsole 是否输出到控制台
func (b *Builder) WriteToConsole(write bool) *Builder {
	b.console = write
	return b
}

// OutputJson 输出 json 格式
func (b *Builder) OutputJson(json bool) *Builder {
	b.outputJson = json
	return b
}

// PrettyJson 是否格式化输出
func (b *Builder) PrettyJson(pretty bool) *Builder {
	b.prettyJson = pretty
	return b
}

// TimestampFormat 设置时间格式
func (b *Builder) TimestampFormat(format string) *Builder {
	b.timestampFormat = format
	return b
}

// SetAppTag 设置应用标志
func (b *Builder) AppTag(appTag string) *Builder {
	b.appTag = appTag
	return b
}

func setNull(logger *logrus.Logger) {
	src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}
	writer := bufio.NewWriter(src)
	logger.SetOutput(writer)
}

func MakeTransfer(cfg *transfers.TransferConfigure) transfers.TransferFn {
	if cfg == nil || cfg.Type == transfers.CONSOLE {
		return transfers.CreateConsoleTransfer(cfg)
	}
	if cfg.Type == transfers.UDP {
		return transfers.CreateUDPTransfer(cfg)
	}
	panic("未实现的转发")
}
