package support_archive

import (
	"fmt"
	"io"

	"github.com/Dynatrace/dynatrace-operator/pkg/logd"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	ctrlzap "sigs.k8s.io/controller-runtime/pkg/log/zap"
)

const (
	supportArchiveLoggerName = "[support-archive]"
)

func newSupportArchiveLogger(out io.Writer) logd.Logger {
	config := zap.NewProductionEncoderConfig()
	config.TimeKey = ""
	config.LevelKey = ""
	config.NameKey = "name"
	config.EncodeTime = zapcore.ISO8601TimeEncoder

	return logd.Logger{
		Logger: ctrlzap.New(ctrlzap.WriteTo(out), ctrlzap.Encoder(zapcore.NewConsoleEncoder(config))).WithName(supportArchiveLoggerName),
	}
}

func logInfof(log logd.Logger, format string, v ...any) {
	log.V(1).Info("foobar")
	log.Info(fmt.Sprintf(format, v...))
}

func logErrorf(log logd.Logger, err error, format string, v ...any) {
	log.Error(err, fmt.Sprintf(format, v...))
}
