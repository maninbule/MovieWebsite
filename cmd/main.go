package main

import (
	"github/MovieWebsite/global/logger"
	"github/MovieWebsite/global/setting"
)

func main() {
	setting.SetupSetting()
	logger.InitLogger()

	logger.Logger.Info("初始化完毕")
}
