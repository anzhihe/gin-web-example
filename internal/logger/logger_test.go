package logger

import (
	"fmt"
	"net/http"
	"testing"
	"thor-backend/internal/setting"

	"github.com/gin-gonic/gin"

	"go.uber.org/zap/zapcore"

	"go.uber.org/zap"
)

var sugarLogger *zap.SugaredLogger

func TestLogger(t *testing.T) {
	// 初始化配置文件
	//if err := setting.Init("../../conf/config.dev.toml"); err != nil {
	//	fmt.Printf("load config failed, err:#{err}\n")
	//	return
	//}
	_ = setting.Init()
	InitLogger(setting.Conf.LogConfig)
	defer sugarLogger.Sync()
	simpleHttpGet("https://chegva.com")
	simpleHttpGet("https://www.mi.com")

	// 初始化logger
	if err := Init(setting.Conf.LogConfig, setting.Conf.Env); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}

	// gin框架日志测试
	SetMode(setting.Conf.Env)

	r := gin.Default()
	// 注册zap相关中间件
	r.Use(GinLogger(), GinRecovery(true))

	r.GET("/test", func(c *gin.Context) {
		// 记录日志中的测试数据
		var (
			username = "anzhihe"
			password = "888888"
			age      = 88
		)
		// 使用zap.x(key, val)将相关字段写入日志
		zap.L().Debug("this is test func", zap.String("username", username), zap.String("password", password), zap.Int("age", age))
		c.String(http.StatusOK, "hello anzhihe!")
	})

	err := r.Run(fmt.Sprintf(":%d", setting.Conf.Port))
	if err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}
}

func InitLogger(cfg *setting.LogConfig) {
	writeSyncer := getLogWriter(cfg.LogPath, cfg.MaxSize, cfg.MaxBackups, cfg.MaxAge)
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	logger := zap.New(core, zap.AddCaller())
	sugarLogger = logger.Sugar()
}

func simpleHttpGet(url string) {
	sugarLogger.Debugf("Trying to hit GET request for %s", url)
	resp, err := http.Get(url)
	if err != nil {
		sugarLogger.Errorf("Error fetching URL %s : Error = %s", url, err)
	} else {
		sugarLogger.Infof("Success! statusCode = %s for URL %s", resp.Status, url)
		resp.Body.Close()
	}
}

func SetMode(mode string) {
	// Gin Mode设置
	switch mode {
	case "online":
		gin.SetMode(gin.ReleaseMode)
	case "dev":
		gin.SetMode(gin.DebugMode)
	case "pre":
		gin.SetMode(gin.TestMode)
	default:
		panic("mode unavailable. (debug, release, test)")
	}
}
