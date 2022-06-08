package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

//使用结构体接收配置文件
type Config struct {
	*AppConf   `mapstructure:"app"`
	*LogConf   `mapstructure:"log"`
	*MySQLConf `mapstructure:"mysql"`
	*RedisConf `mapstructure:"redis"`
}

type AppConf struct {
	Name      string `mapstructure:"name"`
	Mode      string `mapstructure:"mode"`
	Port      int    `mapstructure:"port"`
	StartTime string `mapstructure:"start_time"`
	MachineID int64    `mapstructure:"machine_id"`
}

type LogConf struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"Filename"`
	MaxSize    int    `mapstructure:"MaxSize"`
	MaxAge     int    `mapstructure:"MaxAge"`
	MaxBackups int    `mapstructure:"MaxBackups"`
}

type MySQLConf struct {
	Host         string `mapstructure:"host"`
	Port         string `mapstructure:"port"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	Dbname       string `mapstructure:"dbname"`
	MaxIdleConns int    `mapstructure:"maxidleconns"`
	MaxOpenConns int    `mapstructure:"maxopenconns"`
}

type RedisConf struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Password string `mapstructure:"password"`
	PoolSize int    `mapstructure:"pool_size"`
	Db       int    `mapstructure:"db"`
}

var Conf = new(Config)

func Init() (err error) {
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
	if err = viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.Unmarlshal failed,err:%v\n", err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了...")
		if err = viper.Unmarshal(Conf); err != nil {
			fmt.Printf("viper.Unmarlshal failed,err:%v\n", err)
		}
	})
	return err
}
