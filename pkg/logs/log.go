package logs

import (
	"github.com/seerx/base"
	"github.com/sirupsen/logrus"
)

// Log 默认日志
var Log *logrus.Logger

// InitLog 初始化默认日志
func InitLog(b *Builder) {
	if b == nil {
		Log = NewBuilder().
			Build()
	}
	Log = b.Build()
}

var stack *base.CallStack

func init() {
	stack = base.NewCallStack()
	stack.AddSkipPackage("github.com/sirupsen/logrus")
}

// InitLog 初始化日志组件
//func InitLog(cfg *Configure) {
//	fmt.Printf("初始化日志组件 ...\n")
//
//	Log.Formatter = &logrus.JSONFormatter{
//		TimestampFormat:  time.RFC3339,
//		DisableTimestamp: false,
//		DataKey:          "",
//		FieldMap:         nil,
//		CallerPrettyfier: nil,
//		PrettyPrint:      cfg.PrettyJson,
//	}
//	//Log.Formatter = &logrus.TextFormatter{}
//	setNull(Log)
//	Log.Level = logrus.DebugLevel
//	Log.ReportCaller = true
//	//Log.SetOutput(nil)
//	Log.AddHook(NewTransferHook(cfg.AppTag,
//		Log.Formatter,
//		MakeTransfer(cfg.TransferCfg)))
//}
