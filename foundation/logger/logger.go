package logger

import (
	"fmt"
	"io"
	"path/filepath"

	"golang.org/x/exp/slog"
)

// Logger represents a logger for logging information.
type Logger struct {
	*slog.Logger
}

func New(w io.Writer) *Logger {

	// Convert the file name to just the name.ext when this key/value will
	// be logged.
	f := func(groups []string, a slog.Attr) slog.Attr {
		if a.Key == slog.SourceKey {
			if source, ok := a.Value.Any().(*slog.Source); ok {
				v := fmt.Sprintf("%s:%d", filepath.Base(source.File), source.Line)
				return slog.Attr{Key: "file", Value: slog.StringValue(v)}
			}
		}

		return a
	}

	h := slog.NewJSONHandler(w, &slog.HandlerOptions{AddSource: true, Level: slog.LevelInfo, ReplaceAttr: f})

	return &Logger{
		Logger: slog.New(h),
	}
}
