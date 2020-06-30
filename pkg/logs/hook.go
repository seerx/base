package logs

import (
	"github.com/seerx/base/pkg/logs/transfers"
	"github.com/sirupsen/logrus"
)

// TransferHook 转发钩子
type TransferHook struct {
	tag      string
	fmt      logrus.Formatter
	chs      []chan []byte
	transfer transfers.TransferFn
}

// NewTransferHook 新建转发钩子
func NewTransferHook(tag string, fmt logrus.Formatter, transferFn ...transfers.TransferFn) *TransferHook {
	var chs []chan []byte
	for _, fn := range transferFn {
		ch := make(chan []byte, 500)
		chs = append(chs, ch)
		go fn(ch)
	}
	return &TransferHook{
		tag: tag,
		fmt: fmt,
		chs: chs,
		//transfer: transferFn,
	}
}

// Levels 日志级别
func (t *TransferHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
		logrus.InfoLevel,
		logrus.DebugLevel,
	}
}

// Fire 日志钩子
func (t *TransferHook) Fire(entry *logrus.Entry) error {
	if t.tag != "" {
		entry.Data["app"] = t.tag
	}
	if errObj := entry.Data["error"]; errObj != nil {
		if err, ok := errObj.(error); ok {
			err = stack.WrapErrorSkip(err, 1)
			entry.Data["error"] = err
		}
	}
	data, err := t.fmt.Format(entry)
	if err != nil {
		// 日志格式化错误
		return err
	}
	for _, ch := range t.chs {
		ch <- data
	}
	return nil
}
