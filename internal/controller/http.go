package controller

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"thor-backend/internal/logger"
	"thor-backend/internal/logic"
	"thor-backend/internal/setting"
	log "thor-backend/pkg"
	"time"

	"go.uber.org/zap"

	"github.com/gin-contrib/pprof"

	"github.com/gin-gonic/gin"

	"github.com/toolkits/pkg/cache"
)

type Server struct {
	logic *logic.Logic
	srv   *http.Server
}

func Init(l *logic.Logic) *Server {
	s := &Server{logic: l}

	// 初始化内存缓存
	cache.InitMemoryCache(time.Hour)

	// 设置gin启动模式
	SetMode(setting.Conf.Env)

	// 初始化gin框架
	e := gin.New()
	e.Use(logger.GinLogger(), logger.GinRecovery(true))
	s.initRouter(e)
	pprof.Register(e)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", setting.Conf.Port),
		Handler: e,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	s.srv = srv
	log.Info("server start success.")
	return s
}

func (s *Server) Close() {
	// graceful shutdown
	cxt, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := s.srv.Shutdown(cxt); err != nil {
		log.Error("server shutdown error.", zap.Error(err))
	}
	log.Info("server shutdown...")
}

func GracefulShutdown(s *Server) {
	// init signal
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	for {
		ss := <-sig
		log.Info(fmt.Sprintf("program get a signal %s", ss.String()))
		switch ss {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			s.Close()
			//l.Close()
			//d.Close()
			time.Sleep(2 * time.Second)
			log.Info("program exit.")
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}

func (s *Server) health(c *gin.Context) {
	ResponseSuccess(c, "服务正常")
}

// Gin Mode设置
func SetMode(mode string) {
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
