package logger

import (
	"context"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"io"
	"os"
	"runtime"
	"strings"
	"sync"
)

var once sync.Once

var log zerolog.Logger

func Get() zerolog.Logger {
	once.Do(func() {
		zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
		zerolog.TimeFieldFormat = "2006-01-02T15:04:05.999999Z"

		levelStr := os.Getenv("LOG_LEVEL")
		level := getLogLevel(levelStr)

		var output io.Writer = zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: "2006-01-02T15:04:05.999999Z",
			FormatLevel: func(i interface{}) string {
				return strings.ToUpper(fmt.Sprintf("[%s]", i))
			},
			FormatTimestamp: func(i interface{}) string {
				return fmt.Sprintf("[%s]", i)
			},
			FormatErrFieldName:  func(i interface{}) string { return fmt.Sprintf("*** %s=", i) },
			FormatErrFieldValue: func(i interface{}) string { return fmt.Sprintf("%s ***", i) },
			FieldsOrder: []string{
				"timestamp",
				"level",
				"msg",
				"error",
				"repository.params",
				"service.params",
				"controller.params",
				"stack",
				"go_version",
			},
		}

		env := os.Getenv("APP_ENV")
		if env == "production" {
			output = os.Stderr
		}

		log = zerolog.New(output).
			Level(level).
			With().
			Timestamp().
			Str("go_version", runtime.Version()).
			Logger()
	})

	return log
}

func getLogLevel(logLevelStr string) zerolog.Level {
	switch strings.ToUpper(logLevelStr) {
	case "DEBUG":

		return zerolog.DebugLevel
	case "INFO":

		return zerolog.InfoLevel
	case "WARN":

		return zerolog.WarnLevel
	case "ERROR":
		return zerolog.ErrorLevel
	default:
		return zerolog.InfoLevel
	}
}

type ctxKeyLogger struct{}

func WithLogger(ctx context.Context, logger zerolog.Logger) context.Context {
	return context.WithValue(ctx, ctxKeyLogger{}, logger)
}

func FromContext(ctx context.Context) zerolog.Logger {
	l, ok := ctx.Value(ctxKeyLogger{}).(zerolog.Logger)
	if !ok {
		return zerolog.Nop()
	}
	return l
}
