package pkg

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// ViperInit :Viper 配置文件初始化
func ViperInit() (int, func(), error) {
	// 指定配置文件名称(无扩展名)
	viper.SetConfigName("config")
	// 配置文件类型
	viper.SetConfigType("ini")
	// 添加配置文件路径
	viper.AddConfigPath("./configs/")
	// 读取配置信息
	err := viper.ReadInConfig()
	if err != nil {
		return 0, func() {}, err
	}
	// 开启监控 实现热更新
	viper.WatchConfig()
	// 设置更新回调函数
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("Config Change Success:", in.String())
	})
	return 0, func() {

	}, nil
}
