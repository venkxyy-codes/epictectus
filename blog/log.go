package blog

import (
	"context"
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

var (
	logger Logger
)

func init() {
	logger = &delegatingLogger{}
}

type Logger interface {
	Debug(msg string, keysAndValues ...interface{})
	Info(msg string, keysAndValues ...interface{})
	Warn(msg string, keysAndValues ...interface{})
	Error(err error, msg string, keysAndValues ...interface{})
}

type LogConfig struct {
	Level    string `yaml:"LEVEL" env:"LOG_LEVEL"`
	Format   string `yaml:"FORMAT" env:"LOG_FORMAT"`
	Output   string `yaml:"OUTPUT" env:"LOG_OUTPUT"`
	UnixTime bool   `yaml:"UNIX_TIME" env:"LOG_UNIX_TIME"`
}

func (l *LogConfig) SetDefault() {
	l.Level = "info"
	l.Format = "json"
	l.Output = "stdout"
	l.UnixTime = true
}

func SetupLogger(lcg LogConfig) {
	switch lcg.Output {
	case "stdout":
		l := zerolog.New(os.Stderr).With().Timestamp().Logger()
		logger = &delegatingLogger{log: &l}
	case "console":
		l := log.Output(zerolog.ConsoleWriter{
			Out: os.Stderr,
		})
		logger = &delegatingLogger{log: &l}
	}

	if lcg.UnixTime {
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	}
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
}

func SetLevel(level string) {
	zerolog.SetGlobalLevel(GetLevel(level))
}

func GetLevel(level string) zerolog.Level {
	switch level {
	case "debug", "DEBUG":
		return zerolog.DebugLevel
	case "info", "INFO":
		return zerolog.InfoLevel
	case "warn", "WARN":
		return zerolog.WarnLevel
	case "error", "ERROR":
		return zerolog.ErrorLevel
	case "panic", "PANIC":
		return zerolog.PanicLevel
	default:
		return zerolog.NoLevel
	}
}

func Debug(msg string, keysAndValues ...interface{}) {
	logger.Debug(msg, keysAndValues...)
}

func Info(msg string, keysAndValues ...interface{}) {
	logger.Info(msg, keysAndValues...)
}

func Warn(msg string, keysAndValues ...interface{}) {
	logger.Warn(msg, keysAndValues...)
}

func Error(err error, msg string, keysAndValues ...interface{}) {
	if err == nil {
		logger.Warn(msg, keysAndValues...)
	} else {
		logger.Error(err, msg, keysAndValues...)
	}
}

// GlobalLogger returns a Logger, it is advised to inject and not use directly
func GlobalLogger() Logger {
	return logger
}

type delegatingLogger struct {
	log *zerolog.Logger
}

func (l *delegatingLogger) Debug(msg string, keysAndValues ...interface{}) {
	if l.log != nil {
		f, d := HandleFields(keysAndValues)
		if d != nil {
			l.log.Debug().Msg(*d)
		}
		l.log.Debug().Fields(f).Msg(msg)
	}
}

func (l *delegatingLogger) Info(msg string, keysAndValues ...interface{}) {
	if l.log != nil {
		f, d := HandleFields(keysAndValues)
		if d != nil {
			l.log.Debug().Msg(*d)
		}
		l.log.Info().Fields(f).Msg(msg)
	}
}

func (l *delegatingLogger) Warn(msg string, keysAndValues ...interface{}) {
	if l.log != nil {
		f, d := HandleFields(keysAndValues)
		if d != nil {
			l.log.Debug().Msg(*d)
		}
		l.log.Warn().Fields(f).Msg(msg)
	}
}

func (l *delegatingLogger) Error(err error, msg string, keysAndValues ...interface{}) {
	if l.log != nil {
		f, d := HandleFields(keysAndValues)
		if d != nil {
			l.log.Debug().Msg(*d)
		}
		l.log.Err(err).Fields(f).Msg(msg)
	}
}

func HandleFields(keysAndValues []interface{}) (map[string]interface{}, *string) {
	fields := make(map[string]interface{}, len(keysAndValues)/2)
	for i := 0; i < len(keysAndValues); {
		if i == len(keysAndValues)-1 {
			msg := fmt.Sprintf("odd number of arguments passed as key-value pairs for logging: %+v", keysAndValues[i])
			return fields, &msg
		}

		key, val := keysAndValues[i], keysAndValues[i+1]
		keyStr, isString := key.(string)
		if !isString {
			msg := fmt.Sprintf("non-string key argument passed to logging, ignoring all later arguments: %s", key)
			return fields, &msg
		}

		fields[keyStr] = val
		i += 2
	}
	return fields, nil
}

type ContextValues map[string]interface{}

type KeyType string

const ContextKey KeyType = "BLOG"

func SetValueInContext(ctx context.Context, key string, value interface{}) context.Context {
	if ctx.Value(ContextKey) == nil {
		ctx = context.WithValue(ctx, ContextKey, ContextValues{})
	}
	contextValues := ctx.Value(ContextKey).(ContextValues)
	contextValues[key] = value
	return context.WithValue(ctx, ContextKey, contextValues)
}

func DebugCtx(ctx context.Context, message string, keysAndValues ...interface{}) {
	if ctx.Value(ContextKey) == nil {
		ctx = context.WithValue(ctx, ContextKey, ContextValues{})
	}
	contextValues := ctx.Value(ContextKey).(ContextValues)
	for key, value := range contextValues {
		keysAndValues = append(keysAndValues, key, value)
	}
	Debug(message, keysAndValues...)
}

func InfoCtx(ctx context.Context, message string, keysAndValues ...interface{}) {
	if ctx.Value(ContextKey) == nil {
		ctx = context.WithValue(ctx, ContextKey, ContextValues{})
	}
	contextValues := ctx.Value(ContextKey).(ContextValues)
	for key, value := range contextValues {
		keysAndValues = append(keysAndValues, key, value)
	}
	Info(message, keysAndValues...)
}

func ErrorCtx(ctx context.Context, err error, message string, keysAndValues ...interface{}) {
	if ctx.Value(ContextKey) == nil {
		ctx = context.WithValue(ctx, ContextKey, ContextValues{})
	}
	contextValues := ctx.Value(ContextKey).(ContextValues)
	for key, value := range contextValues {
		keysAndValues = append(keysAndValues, key, value)
	}
	Error(err, message, keysAndValues...)
}
