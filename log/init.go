package log

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func Init(level zapcore.Level) {
	timestamp := time.Now().Format("20060102150405")
	file, _ := os.Create("result-" + timestamp + ".txt")
	fileSyncer := zapcore.AddSync(file)

	consoleSyncer := zapcore.AddSync(os.Stdout)

	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	combinedSyncer := zapcore.NewMultiWriteSyncer(fileSyncer, consoleSyncer)

	core := zapcore.NewCore(encoder, combinedSyncer, level)

	Logger = zap.New(core)
}
