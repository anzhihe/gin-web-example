package main

import (
	"thor-backend/internal/controller"
	"thor-backend/internal/dao"
	"thor-backend/internal/logic"
	"thor-backend/internal/setting"
)

// @title 快速初始化gin web项目
// @version 1.0.1
// @description gin-web-example API接口文档
// @termsOfService https://github.com/anzhihe/gin-web-example
// @contact.name anzhihe
// @contact.url https://chegva.com
// @contact.email anzhihe@foxmail.com

func main() {
	// 加载配置文件
	setting.Init()
	// 加载数据库
	d := dao.Init()
	// 加载应用
	l := logic.Init(d)
	// 启动服务
	s := controller.Init(l)
	// 平滑关闭
	controller.GracefulShutdown(s)
}
