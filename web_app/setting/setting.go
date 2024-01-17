package setting

import (
	"fmt"

	"github.com/fsnotify/fsnotify"

	"github.com/spf13/viper"
)

// Conf 全局变量用来保存程序的所以配置信息
var Conf = new(AppConfig)

type AppConfig struct {
	Name         string `mapstructure:"name"`
	Mode         string `mapstructure:"mode"`
	Version      string `mapstructure:"version"`
	Port         int    `mapstructure:"port"`
	*LogConfig   `mapstructure:"log"`
	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

type LogConfig struct {
	Level    string `mapstructure:"level"`
	Filename string `mapstructure:"filename"`
	MaxSize  int    `mapstructure:"max_siez"`
	MaxAge   int    `mapstructure:"max_age"`
}

type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	Username     string `mapstructure:"username"`
	Password     string `mapstructure:"password"`
	DBname       string `mapstructure:"dbname"`
	Port         int    `mapstructure:"port"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

type RedisConfig struct {
	Host         string `mapstructure:"host"`
	Password     string `mapstructure:"password"`
	Port         int    `mapstructure:"port"`
	DB           int    `mapstructure:"db"`
	PoolSize     int    `mapstructure:"pool_size"`
	MinIdleConns int    `mapstructure:"min_idle_conns"`
}

func Init(filePath string) (err error) {
	//方式1：直接指定配置文件路径（相对路径和绝对路径）
	//相对路径：相对执行的可执行文件的相对路径
	//viper.SetConfigFile("./conf/config.yaml")
	//绝对路径：系统中实际的文件路径
	//viper.SetConfigFile("/User/xx/desktop/bluebell/config.yaml")

	//方式2：指定配置文件名和配置文件的位置 viper自行查找可用的配置文件
	//配置文件名不需要带后缀
	//配置文件位置可配置多个
	//viper.SetConfigFile("config") //指定配置文件名（不带后缀）
	//viper.AddConfigPath(".") //指定查找配置文件的路径（这里使用相对路径）

	viper.SetConfigFile(filePath)
	viper.SetConfigFile("config.yaml")
	//viper.SetConfigName("config") // 配置文件名称(无扩展名)
	//viper.SetConfigType("yaml")   // 如果配置文件的名称中没有扩展名，则需要配置此项(专用于从远程获取配置信息时指定的配置类型信息)
	//viper.AddConfigPath(".")      // 查找配置文件所在的路径(相对绝对路径都可以)
	err = viper.ReadInConfig() // 查找并读取配置文件
	if err != nil {            // 处理读取配置文件的错误
		fmt.Printf("viper.readInConfig failed, err: %v\n", err)
		return
		//panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// 把读取到的配置信息反序列化到 Conf 变量中
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
	}

	//重启后自动更新加载配置
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了...")
		//todo 配置改变
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
		}
	})

	return
	//r := gin.Default()
	//
	//if err := r.Run(); err != nil {
	//	panic(err)
	//}
}
