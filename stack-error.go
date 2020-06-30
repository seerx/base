package base

var callStackInstance *CallStack

func init() {
	callStackInstance = NewCallStack()
}

// type errStack struct {
// 	stackPC []uintptr
// 	raw     error
// }

// ErrorStack 错误堆栈信息
func ErrorStack(err error) error {
	return ErrorStackSkip(err, 0)
	// pcs := make([]uintptr, 32)
	// // skip func StackError invocations
	// count := runtime.Callers(2, pcs)
	// return &errStack{
	// 	raw:     err,
	// 	stackPC: pcs[:count],
	// }
}

// ErrorStackSkip 错误堆栈信息
func ErrorStackSkip(err error, skip int) error {
	return callStackInstance.WrapErrorSkip(err, skip+1)
	// pcs := make([]uintptr, 32)
	// // skip func StackError invocations
	// count := runtime.Callers(2+skip, pcs)
	// return &errStack{
	// 	raw:     err,
	// 	stackPC: pcs[:count],
	// }
}

// github.com/sirupsen/logrus
func (e *errStack) Error() string {
	return e.call.error(e)
}

// func (e *errStack) Error() string {
// 	frames := runtime.CallersFrames(e.stackPC)

// 	var (
// 		f     runtime.Frame
// 		more  bool
// 		index int
// 	)

// 	errString := ""
// 	if e.raw != nil && e.raw.Error() != "" {
// 		errString = e.raw.Error() + "\n"
// 	}

// 	for {
// 		f, more = frames.Next()
// 		if index = strings.Index(f.File, "src"); index != -1 {
// 			// trim GOPATH or GOROOT prifix
// 			f.File = string(f.File[index+4:])
// 		}
// 		errString = fmt.Sprintf("%s%s\n\t%s:%d\n", errString, f.Function, f.File, f.Line)
// 		if !more {
// 			break
// 		}
// 	}
// 	return errString
// }
