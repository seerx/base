package logs

import (
	"bytes"
	"fmt"

	"github.com/sirupsen/logrus"
)

// TextFormatter 自定义日志格式化输出
type TextFormatter struct {
	timeFormat string
}

const (
	formatWithCaller = "[%s] %s %s\n%s:%d\n%s"
	format           = "[%s] %s\n%s"
)

func (t *TextFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var body string
	if entry.HasCaller() {
		body = fmt.Sprintf(formatWithCaller,
			entry.Level.String(),
			entry.Time.Format(t.timeFormat),
			entry.Caller.Func.Name(),
			entry.Caller.File,
			entry.Caller.Line,
			entry.Message)
	} else {
		body = fmt.Sprintf(format,
			entry.Level.String(),
			entry.Time.Format(t.timeFormat),
			entry.Message)
	}
	b := &bytes.Buffer{}
	b.WriteString(body)

	for k, v := range entry.Data {
		b.WriteString(fmt.Sprintf("\n%s:%v", k, v))
	}

	b.WriteString("\n")
	return b.Bytes(), nil
}
