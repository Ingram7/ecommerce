.PHONY: test
#NOW = $(shell date -u '+%Y%m%d%I%M%S')

RELEASE_VERSION = 1.0.0

APP 			= ecommerce
SERVER_BIN  	= ./cmd/bin/${APP}
# git 提交次数
GIT_COUNT 		= $(shell git rev-list --all --count)
# 当前提交版本
GIT_HASH        = $(shell git rev-parse --short HEAD)
RELEASE_TAG     = $(RELEASE_VERSION).$(GIT_COUNT).$(GIT_HASH)

#-w 去掉DWARF调试信息，得到的程序不能用gdb调试
#-s 去掉符号表,panic时候的stack trace就没有任何文件名/行号信息
#-X 设置包中的变量值
build:
	@go build -ldflags "-w -s -X main.VERSION=$(RELEASE_TAG)" -o $(SERVER_BIN) ./cmd/

start:
	go run -ldflags "-X main.VERSION=$(RELEASE_TAG) -X $(APP)/app.Path=./cmd/" ./cmd/main.go ecommerce

startServer:
	nohup go run -ldflags "-X main.VERSION=$(RELEASE_TAG) -X $(APP)/app.Path=./cmd/" ./cmd/main.go ecommerce &

test:
	go test -v ./...  -ldflags "-X $(APP)/app.Path=../cmd/"

swag:
	swag init --parseDependency --generalInfo ./cmd/main.go --output ./interfaces/swagger

ent:
	go generate ./repository/ent

stop:
	ps -ef |grep ecommerce|grep -v 'grep' |awk '{print $2}'|xargs kill -9