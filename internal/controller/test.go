package controller

import (
	log "thor-backend/pkg"

	"github.com/gin-gonic/gin"
)

// ServeTest router测试接口
// @Summary router测试接口
// @Description 测试router访问是否正常
// @Tags Router测试相关接口
// @Produce application/json
// @Param name body string true "测试名称" minlength(3) maxlength(100)
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {string} string "请求成功"
// @Failure 400 "请求错误"
// @Failure 500 "内部繁忙"
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
