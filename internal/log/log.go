package log

import (
	"errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"syscall"
)

const logLayout = "2006-01-02 15:04:05.000"

func Init(level string) (logger *zap.Logger, err error) {
	Level := zap.NewAtomicLevel()
	if err = Level.UnmarshalText([]byte(level)); err != nil {
		return nil, err
	}

	logger, err = zap.Config{
		Level:             Level,
		Development:       false,
		Encoding:          "json",
		DisableStacktrace: true,
		DisableCaller:     true,
		OutputPaths:       []string{"stdout"},
		ErrorOutputPaths:  []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "ts",
			EncodeTime:     zapcore.TimeEncoderOfLayout(logLayout),
			EncodeDuration: zapcore.StringDurationEncoder,

			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,

			NameKey:     "key",
			FunctionKey: zapcore.OmitKey,

			MessageKey: "msg",
			LineEnding: zapcore.DefaultLineEnding,
		},
	}.Build()

	return logger, err
}

func Close(logger *zap.Logger) error {
	if err := logger.Sync(); err != nil && !errors.Is(err, syscall.ENOTTY) {
		return err
	}
	return nil
}
