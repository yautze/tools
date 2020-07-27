package logger

import (
	"testing"
	"time"
)

func Test_Logrus(t *testing.T) {
	NewLogrus()
	SetServiceInfo("foreman")
	WithFields(Fields{
		"msg": time.Now().Unix(),
	}).Info("12")
}