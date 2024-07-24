package sl

import (
	"golang.org/x/exp/slog"
)

// по-другому назвать пакет??? -> почему не засунуть сразу с логгером рядом?

func Err(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}
