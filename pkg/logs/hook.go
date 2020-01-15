package logs

import (
	"github.com/seerx/base/pkg/logs/transfers"
	"github.com/sirupsen/logrus"
)

type TransferHook struct {
	tag      string
	fmt      logrus.Formatter
	chs      []chan []byte
	transfer transfers.TransferFn
}

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

func (t *TransferHook) Fire(entry *logrus.Entry) error {
	if t.tag != "" {
		entry.Data["app"] = t.tag
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
