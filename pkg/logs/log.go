package logs

import (
	"github.com/sirupsen/logrus"
)

//type Configure struct {
//	AppTag      string
//	PrettyJson  bool
//	TransferCfg *transfers.TransferConfigure
//}

var Log *logrus.Logger

func InitLog(b *Builder) {
	if b == nil {
		Log = NewBuilder().
			Build()
	}
	Log = b.Build()
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
