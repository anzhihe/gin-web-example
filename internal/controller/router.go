package controller

import (
	_ "thor-backend/docs" // 导入swag生成的docs
	"thor-backend/pkg/app"

	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/gin-gonic/gin"
)

func (s *Server) initRouter(e *gin.Engine) {
	// health check
	e.GET("/health", s.health)

	// init swagger
	e.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	v1 := e.Group("/api/v1")
	{
		// 测试接口
		v1.GET("/test", s.ServeTest)
	}

	e.NoRoute(func(c *gin.Context) {
		app.ResponseError(c, app.CodeNotFound)
	})

}