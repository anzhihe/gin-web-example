package controller

// 分页处理

import (
	"thor-backend/internal/setting"
	"thor-backend/pkg/convert"

	"github.com/gin-gonic/gin"
)

func GetPage(c *gin.Context) int {
	page := convert.StrTo(c.Query("page")).MustInt()
	if page <= 0 {
		return 1
	}

	return page
}

func GetPageSize(c *gin.Context) int {
	pageSize := convert.StrTo(c.Query("page_size")).MustInt()
	if pageSize <= 0 {
		return setting.Conf.DefaultPageSize
	}
	if pageSize > setting.Conf.MaxPageSize {
		return setting.Conf.MaxPageSize
	}

	return pageSize
}

func PageOffset(page, pageSize int) int {
	result := 0
	if page > 0 {
		result = (page - 1) * pageSize
	}

	return result
}
