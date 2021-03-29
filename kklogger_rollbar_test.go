package kkrollbar

import (
	"testing"
	"time"

	"github.com/kklab-com/goth-kklogger"
)

func TestKKLoggerRollbarHook(t *testing.T) {
	hook := KKLoggerRollbarHook{
		Token: "",
		Level: kklogger.DebugLevel,
	}

	kklogger.AsyncWrite = false
	kklogger.SetLoggerHooks([]kklogger.LoggerHook{&hook})
	kklogger.SetLogLevel("DEBUG")
	kklogger.DebugJ("djsType", "jsData")
	time.Sleep(time.Second * 5)
}
