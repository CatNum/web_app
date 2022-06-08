package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Init1() (err error) {
	//viper.SetConfigFile("config.yaml")   //当配置文件有多个比如config.yaml、config.json，使用该函数指定文件
	viper.SetConfigName("config") // 配置文件名称(不需要带后缀)
	// 如果配置文件的名称中没有扩展名，则需要配置此项(专用于从远程获取配置信息时指定配置文件类型的，从配置文件中加载配置则不生效)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")   // 查找配置文件所在的路径
	err = viper.ReadInConfig() // 查找并读取配置文件
	if err != nil {            // 处理读取配置文件的错误
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
		return
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了...")
	})
	return err
}
