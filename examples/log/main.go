package main

import (
	"errors"
	"time"

	"github.com/seerx/base"
	"github.com/seerx/base/pkg/logs"
)

func test() {
	logs.WithError(errors.New("111")).Info("123")
}

func main() {
	logs.InitLog(logs.NewBuilder().
		TimestampFormat(base.TFDatetimeMilli))
	// fmt.Fprint(os.Stderr, "111111")
	test()
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
