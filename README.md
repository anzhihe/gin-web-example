## Gin Web Example

An example of gin web project, using Gin、Gorm、MySQL、Redis etc.

[简体中文](./README_ZH.md) | English

## How to run

### Required
- MySQL
- Redis

### Ready
1. Start MySQL & Redis
2. Create three databases: ordin、inspection、risk
```bash
mysql -uroot -prootroot -e "CREATE DATABASE IF NOT EXISTS ordin DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;CREATE DATABASE IF NOT EXISTS inspection DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;CREATE DATABASE IF NOT EXISTS risk DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;"
```

### Conf
You should modify `conf/config.dev.toml`

```bash
[database]
Type = mysql
User = root
Password = rootroot
Host = 127.0.0.1:3306
ODBName =  ordin 
IDBName = inspection 
RDBName = risk

[redis]
Host = 127.0.0.1:6379
Password = ""
db = 0
pool_size = 100
```

## Run

Specify the configuration file `go run main.go -c conf/config.dev.toml`
```bash
$ git clone https://github.com/anzhihe/gin-web-example.git

$ cd gin-web-example

$ go run main.go or make run or air
```
Project information and existing API
```bash
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /health                   --> thor-backend/internal/controller.(*Server).health-fm (3 handlers)
[GIN-debug] GET    /swagger/*any             --> github.com/swaggo/gin-swagger.CustomWrapHandler.func1 (3 handlers)
[GIN-debug] GET    /api/v1/test              --> thor-backend/internal/controller.(*Server).GetServeTest-fm (3 handlers)
[GIN-debug] GET    /debug/pprof/             --> github.com/gin-contrib/pprof.pprofHandler.func1 (3 handlers)
......
Listening port is 8080
```

### Swagger doc

- http://127.0.0.1:8080/swagger/index.html

## Directory 

```bash
thor-backend
├── Makefile        // 编译打包
├── README.md       // 使用说明
├── conf            // 配置文件
│   ├── config.dev.toml
│   ├── config.online.toml
│   └── config.pre.toml
├── docs            // 文档目录
├── internal    
│   ├── controller  // 路由入口
│   ├── dao         // 数据操作
│   ├── logger      // 日志封装
│   ├── logic       // 业务逻辑
│   ├── middleware  // 中间件层
│   ├── model       // 模型定义
│   └── setting     // 配置定义
├── main.go         // 程序入口
├── .air.conf       // 热重启配置
├── pkg             // 外部通用包
├── scripts         // 项目相关脚本
│   └── sql
└── vendor          // go mod依赖库
```

## Features

- [Github RESTful API](https://docs.github.com/cn/rest)
- [Gin](https://github.com/gin-gonic/gin)
- [Gorm](https://gorm.io)
- [Swagger](https://github.com/swaggo/gin-swagger)
- [zap](https://github.com/uber-go/zap)
- [viper](https://github.com/spf13/viper)
- [air](https://github.com/cosmtrek/air)
- [MySQL](https://www.mysql.com/)
- [Redis](https://github.com/redis/redis)
- App configurable(Toml)
- Graceful restart or stop
- Live reload for Server
