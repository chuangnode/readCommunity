package initialize

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"readCommunity/global"
)

// InitLogger : initialize logger
func InitLogger() (err error) {
	var (
		filename   = viper.GetString("Log.Filename")
		MaxSize    = viper.GetInt("Log.MaxSize")
		MaxBackups = viper.GetInt("Log.MaxBackups")
		MaxAge     = viper.GetInt("Log.MaxAge")
		Level      = viper.GetString("Log.Level")
	)
	fmt.Println("chuang:>>>>filename:", filename)
	fmt.Println(viper.GetString("AppSecret"))
	fmt.Println(Level)
	fmt.Println("**************")
	writeSyncer := getLogWriter(filename, MaxSize, MaxBackups, MaxAge)
	encoder := getEncoder()
	var l = new(zapcore.Level)
	err = l.UnmarshalText([]byte(Level))
	if err != nil {
		return
	}
	core := zapcore.NewCore(encoder, writeSyncer, l)

	global.Lg = zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(global.Lg)
	return
}
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
	}
	return zapcore.AddSync(lumberJackLogger)
}
