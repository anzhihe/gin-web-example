package controller

import (
	log "thor-backend/pkg"

	"github.com/gin-gonic/gin"
)

// ServeTest router测试接口
// @Summary router测试接口
// @Description 测试router访问是否正常
// @Tags Router测试相关接口
// @Accept application/json
// @Produce application/json
// @Success 200
// @Router /api/v1/test [get]
func (s *Server) ServeTest(c *gin.Context) {
	// 获取数据
	data := s.logic.ServeTest()
	count := 100
	log.Info("测试成功")
	log.Error("测试失败")
	// 返回响应
	//ResponseSuccess(c, data)

	ResponseList(c, data, count) // 返回分页数据
}
