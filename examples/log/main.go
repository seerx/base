package main

import (
	"errors"
	"time"

	"github.com/seerx/base"
	"github.com/seerx/base/pkg/logs"
)

func main() {
	logs.InitLog(logs.NewBuilder().
		//AppTag("ycjk").
		TimestampFormat(base.TFDatetimeMilli).
		ReportCaller(false))
	// fmt.Fprint(os.Stderr, "111111")
	logs.Log.WithError(errors.New("111")).Info("123")
	// 	main.main
	//     /Users/dotjava/workspace/go-projects/base/examples/log/main.go:17
	//   runtime.main
	//     runtime/proc.go:203
	//   runtime.goexit
	//     runtime/asm_amd64.s:1373

	time.Sleep(1 * time.Second)
	// err := base.ErrorStackSkip(errors.New("1234456789876543"), 1)
	// fmt.Println(err.Error())
}
