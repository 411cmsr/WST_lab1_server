package logging

import (
	"WST_lab1_server/config"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func Init() {
	var level zapcore.Level
	//Уровень логирования из файла конфигурации
	switch config.GeneralServerSetting.LogLevel {
	case "fatal":
		level = zapcore.FatalLevel
	case "error":
		level = zapcore.ErrorLevel
	case "warn":
		level = zapcore.WarnLevel
	case "info":
		level = zapcore.InfoLevel
	case "debug":
		level = zapcore.DebugLevel
	default:
		level = zapcore.DebugLevel
	}

	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(config)

	core := zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), level)
	Logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	defer Logger.Sync()
}
