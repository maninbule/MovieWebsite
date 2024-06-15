package main

import (
	"github/MovieWebsite/global/db"
	"github/MovieWebsite/global/logger"
	"github/MovieWebsite/global/setting"
)

func main() {
	setting.SetupSetting()
	mylogger.InitLogger()
	db.InitDBEngine()
	mylogger.Logger.Info(nil, "程序启动成功")
}
