package controller

import (
	"fmt"
	"testing"
	"thor-backend/internal/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func TestCode(t *testing.T) {

	r := gin.Default()
	// 注册zap相关中间件
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.GET("/test", func(c *gin.Context) {
		// 记录日志中的测试数据
		var (
			username = "anzhihe"
			password = "888888"
			age      = 88
		)
		// 使用zap.x(key, val)将相关字段写入日志
		zap.L().Debug("this is test func", zap.String("username", username), zap.String("password", password), zap.Int("age", age))
		//c.String(http.StatusOK, "hello anzhihe!")
		ResponseSuccess(c, "xxxxx")
		//ResponseError(c, CodeInvalidToken)
		//ResponseError(c, 10000)
	})

	err := r.Run(fmt.Sprintf(":%d", 8080))
	if err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}
}
