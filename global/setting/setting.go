package setting

import (
	"github.com/spf13/viper"
)

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
