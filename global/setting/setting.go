package global

import (
	"fmt"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper"
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

// 使用viper从配置文件中读取各个组件的配置信息

type Setting struct {
	vp *viper.Viper
}

func NewSetting() (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath("global/setting/")
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return &Setting{vp}, nil
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	return s.vp.UnmarshalKey(k, v)
}
