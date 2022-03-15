package setting

import (
	"flag"
	"fmt"

	//"github.com/BurntSushi/toml"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var (
	Conf     = new(Config)
	confPath string
	err      error
)

// 初始化热加载配置文件
func Init() {
	// &confPath 接收用户命令行中输入的 -c 后面的参数值指定配置文件，"./conf/config.dev.toml" 默认配置文件路径
	flag.StringVar(&confPath, "c", "./conf/config.dev.toml", "need config file.eg: thor -c ./conf/config.[dev|pre|online].toml")
	flag.Parse()

	// 读取配置信息
	viper.SetConfigFile(confPath)
	if err := viper.ReadInConfig(); err != nil {
		// 读取配置信息失败
		fmt.Printf("viper.ReadInConfig failed！\n")
		panic(err)
	}

	// 把读取到的配置信息反序列化到 Conf 变量中
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.Unmarshal failed！\n")
		panic(err)
	}

	// 监听配置文件
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Printf("配置文件已被修改。")
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("viper.Unmarshal failed！\n")
			panic(err)
		}
	})
}

// 配置文件结构体
type Config struct {
	*RunConfig `mapstructure:"run"`
	*AppConfig `mapstructure:"app"`
	*LogConfig `mapstructure:"log"`
	ODB        *DBConfig    `mapstructure:"odb"` // 作业平台数据库
	IDB        *DBConfig    `mapstructure:"idb"` // 巡检平台数据库
	RDB        *DBConfig    `mapstructure:"rdb"` // 预警平台数据库
	RDS        *RedisConfig `mapstructure:"redis"`
}

// 服务启动配置
type RunConfig struct {
	Name    string `mapstructure:"name"`
	Port    int    `mapstructure:"port"`
	Env     string `mapstructure:"env"`
	Version string `mapstructure:"version"`
}

// 应用配置
type AppConfig struct {
	DefaultPageSize   int    `mapstructure:"default_page_size"`
	MaxPageSize       int    `mapstructure:"max_page_size"`
	UploadSavePath    string `mapstructure:"upload_save_path"`
	UploadServerUrl   string `mapstructure:"upload_server_url"`
	UploadFileMaxSize string `mapstructure:"upload_file_max_size"`
}

// 数据库配置
type DBConfig struct {
	Addr         string `mapstructure:"addr"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
	Debug        bool   `mapstructure:"debug"`
}

type RedisConfig struct {
	Addr     string `mapstructure:"addr"`
	Password string `mapstructure:"password"`
	PoolSize int    `mapstructure:"pool_size"`
}

// 日志配置
type LogConfig struct {
	Level      string `mapstructure:"level"`
	AccessLog  string `mapstructure:"access_log"`
	ErrorLog   string `mapstructure:"error_log"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
	Compress   bool   `mapstructure:"compress"`
}
