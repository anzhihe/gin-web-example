package setting

import (
	"fmt"
	"testing"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/k0kubun/pp/v3"
	"github.com/spf13/viper"
)

func ParseConfig(filePath string) (err error) {

	viper.SetConfigFile(filePath) // 读取配置文件信息路径

	err = viper.ReadInConfig() // 读取配置信息
	if err != nil {
		// 读取配置信息失败
		fmt.Printf("viper.ReadInConfig failed, err:#{err}\n")
		return
	}

	// 把读取到的配置信息反序列化到 Conf 变量中
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.Unmarshal failed, err:#{err}\n")
	}

	// 监听配置文件
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Printf("配置文件已被修改。")
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("viper.Unmarshal failed, err:#{err}\n")
		}
	})
	return
}

// 解析TOML文件
func TestConfigTOML(t *testing.T) {
	if err := ParseConfig("../../conf/config.dev.toml"); err != nil {
		t.Error(err)
		return
	}

	pp.Print(Conf)
}

// 测试更新配置文件热加载
func TestWatchConfig(t *testing.T) {
	if err := ParseConfig("../../conf/config.dev.toml"); err != nil {
		t.Error(err)
		return
	}

	for i := 0; i < 30; i++ {
		pp.Print(Conf)
		time.Sleep(time.Second)
	}
}
