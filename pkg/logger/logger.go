package logger

import (
	"context"
	"fmt"
	"github.com/go-kit/log"
	"io"
	"time"
)

// Logger specifies logging API.
type Logger interface {
	// Debug logs any object in JSON format on debug level.
	Debug(string)
	// Info logs any object in JSON format on info level.
	Info(string)
	// Warn logs any object in JSON format on warning level.
	Warn(string)
	// Error logs any object in JSON format on error level.
	Error(string)

	Debugf(string, ...interface{})
	Infof(string, ...interface{})
	Warnf(string, ...interface{})
	Errorf(string, ...interface{})

	Log(keyvals ...interface{}) error

	Level() string

	LogI(context.Context, string, ...interface{})
	LogW(context.Context, string, ...interface{})
	LogE(context.Context, string, ...interface{})
	LogD(context.Context, string, ...interface{})
	LogS(context.Context, string, ...interface{})
	LogRS(data string)
	LogRE(data string)
}

var _ Logger = (*logger)(nil)

type logger struct {
	kitLogger log.Logger
	level     Level
}

func (l *logger) Debugf(s string, i ...interface{}) {
	if Debug.isAllowed(l.level) {
		l.kitLogger.Log("l", Debug.String(), "m", fmt.Sprintf(s, i...))
	}
}
func (l *logger) Infof(s string, i ...interface{}) {
	if Info.isAllowed(l.level) {
		l.kitLogger.Log("l", Info.String(), "m", fmt.Sprintf(s, i...))
	}
}
func (l *logger) Warnf(s string, i ...interface{}) {
	if Warn.isAllowed(l.level) {
		l.kitLogger.Log("l", Warn.String(), "m", fmt.Sprintf(s, i...))
	}
}
func (l *logger) Errorf(s string, i ...interface{}) {
	if Error.isAllowed(l.level) {
		l.kitLogger.Log("l", Error.String(), "m", fmt.Sprintf(s, i...))
	}
}

// New returns wrapped go kit logger.
func New(out io.Writer, levelText string) (Logger, error) {
	var level Level
	err := level.UnmarshalText(levelText)
	if err != nil {
		return nil, fmt.Errorf(`{"l":"error","m":"%s: %s","ts":"%s"}`, err, levelText, time.DateTime)
	}
	l := log.NewLogfmtLogger(log.NewSyncWriter(out))
	l = log.With(l, "ts", log.DefaultTimestampUTC)
	return &logger{l, level}, err
}

func (l *logger) Log(keyvals ...interface{}) error {
	return l.kitLogger.Log(keyvals)
}

func (l *logger) Debug(msg string) {
	if Debug.isAllowed(l.level) {
		l.kitLogger.Log("l", Debug.String(), "m", msg)
	}
}

func (l *logger) Info(msg string) {
	if Info.isAllowed(l.level) {
		l.kitLogger.Log("l", Info.String(), "m", msg)
	}
}

func (l *logger) Warn(msg string) {
	if Warn.isAllowed(l.level) {
		l.kitLogger.Log("l", Warn.String(), "m", msg)
	}
}

func (l *logger) Error(msg string) {
	if Error.isAllowed(l.level) {
		l.kitLogger.Log("l", Error.String(), "m", msg)
	}
}

func (l *logger) LogE(ctx context.Context, s string, i ...interface{}) {
	if Error.isAllowed(l.level) {
		requestId, ok := ctx.Value(RequestId).(string)
		if !ok {
			requestId = ""
		}
		l.kitLogger.Log("log", Error.String(), "i", requestId, "m", fmt.Sprintf(s, i...))
	}
}
func (l *logger) LogW(ctx context.Context, s string, i ...interface{}) {
	if Warn.isAllowed(l.level) {
		requestId, ok := ctx.Value(RequestId).(string)
		if !ok {
			requestId = ""
		}
		l.kitLogger.Log("log", Warn.String(), "i", requestId, "m", fmt.Sprintf(s, i...))
	}
}
func (l *logger) LogI(ctx context.Context, s string, i ...interface{}) {
	if Info.isAllowed(l.level) {
		requestId, ok := ctx.Value(RequestId).(string)
		if !ok {
			requestId = ""
		}
		l.kitLogger.Log("log", Info.String(), "i", requestId, "m", fmt.Sprintf(s, i...))
	}
}

const RequestId = "requestId"

func (l *logger) LogD(ctx context.Context, s string, i ...interface{}) {
	if Debug.isAllowed(l.level) {
		requestId, ok := ctx.Value(RequestId).(string)
		if !ok {
			requestId = ""
		}
		l.kitLogger.Log("log", Debug.String(), "i", requestId, "m", fmt.Sprintf(s, i...))
	}
}
func (l *logger) LogS(ctx context.Context, s string, i ...interface{}) {
	if System.isAllowed(l.level) {
		requestId, ok := ctx.Value(RequestId).(string)
		if !ok {
			requestId = ""
		}
		l.kitLogger.Log("log", System.String(), "i", requestId, "m", fmt.Sprintf(s, i...))
	}
}

func (l *logger) LogRS(data string) {
	if System.isAllowed(l.level) {
		fmt.Println("LOG|" + time.Now().Format("01/02 15:04:05") + "|" + Debug.String() + "|" + data)
	}
}

func (l *logger) LogRE(data string) {
	if System.isAllowed(l.level) {
		fmt.Println("LOG|" + time.Now().Format("01/02 15:04:05") + "|" + Debug.String() + "|" + data)
	}
}

func (l *logger) Level() string {
	return l.level.String()
}