package setting

import (
	"fmt"
	"time"
)

var (
	ServerSetting   ServerSettings
	AppSetting      AppSettings
	DatabaseSetting DatabaseSettings
)

// SetupSetting 从配置文件中进行初始化
func SetupSetting() {
	setting, err := NewSetting()
	if err != nil {
		panic("viper配置文件没有加载成功")
	}
	err = setting.ReadSection("Server", &ServerSetting)
	if err != nil {
		panic("viper配置文件Server section没有读取成功")
	}
	err = setting.ReadSection("App", &AppSetting)
	if err != nil {
		panic("viper配置文件App section没有读取成功")
	}
	err = setting.ReadSection("Database", &DatabaseSetting)
	if err != nil {
		panic("viper配置文件Database section没有读取成功")
	}
	ServerSetting.ReadTimeout *= time.Second
	ServerSetting.WriteTimeout *= time.Second

	fmt.Printf("%#v\n", ServerSetting)
	fmt.Printf("%#v\n", AppSetting)
	fmt.Printf("%#v\n", DatabaseSetting)

}
