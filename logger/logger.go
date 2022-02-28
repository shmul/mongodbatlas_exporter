package logger

import (
	"fmt"
	"os"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	kl "github.com/go-kit/kit/log/logrus"
	"github.com/sirupsen/logrus"
)

func toLevel(l string) level.Option {
	switch l {
	case "error":
		return level.AllowError()
	case "warn":
		return level.AllowWarn()
	case "info":
		return level.AllowInfo()
	case "debug":
		return level.AllowDebug()
	}
	return nil
}
func Create(l string) (logger log.Logger, err error) {
	lvl := toLevel(l)
	if lvl == nil {
		return nil, fmt.Errorf("unrecognized log level: %v", l)
	}
	logger = log.NewJSONLogger(log.NewSyncWriter(os.Stderr))
	logger = level.NewFilter(logger, lvl)
	logger = log.With(logger, "timestamp", log.DefaultTimestampUTC)

	return logger, nil
}

func CreateUsingLogrus(l string, lgr logrus.FieldLogger) (logger log.Logger, err error) {
	lvl := toLevel(l)
	if lvl == nil {
		return nil, fmt.Errorf("unrecognized log level: %v", l)
	}

	return kl.NewLogger(lgr), nil
}
