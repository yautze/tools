package logger

import "github.com/sirupsen/logrus"

// Fields -
type Fields = logrus.Fields

// Level -
type Level = logrus.Level

// log level
const (
	TraceLevel = logrus.TraceLevel
	DebugLevel = logrus.DebugLevel
	InfoLevel  = logrus.InfoLevel
	ErrorLevel = logrus.ErrorLevel
	WarnLevel  = logrus.WarnLevel
	FatalLevel = logrus.FatalLevel
	PainLevel  = logrus.PanicLevel
)

// FieldHook - 用來增加Name 欄位
type fieldHook struct {
	Service string
	Host    string
}

// Fire -
func (hook *fieldHook) Fire(entry *logrus.Entry) error {
	entry.Data["service"] = hook.Service
	entry.Data["hostname"] = hook.Host

	return nil
}

// Levels -
func (hook *fieldHook) Levels() []logrus.Level {
	return logrus.AllLevels
}