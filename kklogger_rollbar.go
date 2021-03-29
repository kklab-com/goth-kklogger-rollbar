package kkrollbar

import (
	"fmt"
	"sync"

	"github.com/kklab-com/goth-kklogger"
	"github.com/rollbar/rollbar-go"
)

type KKLoggerRollbarHook struct {
	Token       string
	Environment string
	CodeVersion string
	ServerHost  string
	ServerRoot  string
	Level       kklogger.Level
	once        sync.Once
}

func (h *KKLoggerRollbarHook) _Init() {
	h.once.Do(func() {
		rollbar.SetToken(h.Token)
		rollbar.SetEnvironment(h.Environment)
		rollbar.SetCodeVersion(h.CodeVersion)
		rollbar.SetServerHost(h.ServerHost)
		rollbar.SetServerRoot(h.ServerRoot)
	})
}

func (h *KKLoggerRollbarHook) LogString(args ...interface{}) string {
	if args == nil {
		return ""
	}

	args = args[0].([]interface{})
	argl := len(args)

	if argl == 1 {
		switch tp := args[0].(type) {
		case string:
			return tp
		}
	} else if argl > 1 {
		switch tp := args[0].(type) {
		case string:
			pargs := args[1:]
			return fmt.Sprintf(tp, pargs...)
		}
	}

	return fmt.Sprint(args...)
}

func (h *KKLoggerRollbarHook) Trace(args ...interface{}) {
}

func (h *KKLoggerRollbarHook) Debug(args ...interface{}) {
	if h.Level < kklogger.DebugLevel {
		return
	}

	h._Init()
	rollbar.Debug(h.LogString(args...))
}

func (h *KKLoggerRollbarHook) Info(args ...interface{}) {
	if h.Level < kklogger.InfoLevel {
		return
	}

	h._Init()
	rollbar.Info(h.LogString(args...))
}

func (h *KKLoggerRollbarHook) Warn(args ...interface{}) {
	if h.Level < kklogger.WarnLevel {
		return
	}

	h._Init()
	rollbar.Warning(h.LogString(args...))
}

func (h *KKLoggerRollbarHook) Error(args ...interface{}) {
	if h.Level < kklogger.ErrorLevel {
		return
	}

	h._Init()
	rollbar.Error(h.LogString(args...))
}
