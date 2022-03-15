package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 返回数据定义
type ResponseData struct {
	Code ResCode     `json:"code"`           // 业务响应状态码
	Msg  interface{} `json:"msg"`            // 提示信息
	Data interface{} `json:"data,omitempty"` // 返回数据
}

// 分页定义
type Pager struct {
	Page      int `json:"page"`       // 页码
	PageSize  int `json:"page_size"`  // 每页数量
	TotalRows int `json:"total_rows"` // 总行数
}

func ResponseError(c *gin.Context, code ResCode) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	})
}

func ResponseErrorWithMsg(c *gin.Context, code ResCode, msg interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: CodeSuccess,
		Msg:  CodeSuccess.Msg(),
		Data: data,
	})
}

// 返回分页数据
func ResponseList(c *gin.Context, list interface{}, totalRows int) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: CodeSuccess,
		Msg:  CodeSuccess.Msg(),
		Data: map[string]interface{}{
			"List": list,
			"Pager": Pager{
				Page:      GetPage(c),
				PageSize:  GetPageSize(c),
				TotalRows: totalRows,
			},
		},
	})
}
