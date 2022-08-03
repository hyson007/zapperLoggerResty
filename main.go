package main

import (
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func logInit(d bool, f *os.File) *zap.SugaredLogger {

	pe := zap.NewDevelopmentEncoderConfig()
	fileEncoder := zapcore.NewJSONEncoder(pe)
	pe.EncodeTime = zapcore.ISO8601TimeEncoder // The encoder can be customized for each output
	consoleEncoder := zapcore.NewConsoleEncoder(pe)

	// fileEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	// consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())

	level := zap.InfoLevel
	if d {
		level = zap.DebugLevel
	}

	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./test.log", // Log name
		MaxSize:    1,            // File content size, MB
		MaxBackups: 5,            // Maximum number of old files retained
		MaxAge:     30,           // Maximum number of days to keep old files
		Compress:   false,        // Is the file compressed
	}

	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, zapcore.AddSync(lumberJackLogger), level),
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), level),
	)

	l := zap.New(core, zap.AddCaller()) // Creating the logger

	return l.Sugar()
}

func main() {
	f, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	log := logInit(false, f)
	log.Info("Hello, World!")

	L, _ := zap.NewDevelopment()
	log = L.Sugar()
	// log = L.Sugar()
	log.Info("Hello, World!")
}
