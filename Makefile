.PHONY: all build run gotool clean help

# 根据不同平台编译生成不同格式的二进制包
ifeq ($(shell uname),Darwin)
 PLATFORM="darwin"
else
 ifeq ($(OS),Windows_NT)
  PLATFORM="windows"
 else
  PLATFORM="linux"
 endif
endif

BINARY="thor"

LOGDIR = "log"

all: gotool build

build:
	CGO_ENABLED=0 GOOS=${PLATFORM} GOARCH=amd64 go build -o ${BINARY}

run:
	@go run ./main.go conf/config.dev.toml

gotool:
	go fmt ./
	go vet ./

clean:
	@if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi
	@if [ -d ${LOGDIR} ] ; then rm -rf ${LOGDIR} ; fi

help:
	@echo "make - 格式化 Go 代码, 并编译生成二进制文件"
	@echo "make build - 编译 Go 代码, 生成二进制文件"
	@echo "make run - 直接运行 Go 代码"
	@echo "make clean - 移除二进制文件和日志目录"
	@echo "make gotool - 运行 Go 工具 'fmt' and 'vet'"