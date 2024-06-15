package logger

import (
	"github/MovieWebsite/global/setting"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
)

var (
	Logger *LoggerImp
)

func InitLogger() {
	fileName := setting.AppSetting.LogSavePath + "/" + setting.AppSetting.LogFileName + setting.AppSetting.LogFileExt
	Logger = NewLogger(&lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)

}
