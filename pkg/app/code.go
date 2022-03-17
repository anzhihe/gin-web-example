package app

// 自定义返回状态码

type ResCode int64

const (
	CodeSuccess ResCode = 200
	CodeCreated ResCode = 201
	CodeDeleted ResCode = 204

	CodeInvalidParam ResCode = 400
	CodeUnAuthorized ResCode = 401
	CodeForbidden    ResCode = 403
	CodeNotFound     ResCode = 404
	CodeInvalidToken ResCode = 419

	CodeServerBusy ResCode = 500
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:      "请求成功",
	CodeCreated:      "创建成功",
	CodeDeleted:      "删除成功",
	CodeInvalidParam: "请求参数错误",
	CodeUnAuthorized: "未授权",
	CodeForbidden:    "禁止访问",
	CodeNotFound:     "请求资源不存在",
	CodeInvalidToken: "无效的token",
	CodeServerBusy:   "服务繁忙",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}
