package app

import (
	"ecommerce/pkg/app"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type engine struct {
	conf   *config
	mode   *app.Mode

	server *server
}

func NewEngine() *engine {
	e := new(engine)
	e.conf = loadConfig()
	e.conf.print()
	return e
}

func (e *engine) WarmUp() {
	e.mode = app.NewMode(e.conf.Mode)
	e.server = newServer(e.conf, e.mode)
}

func (e *engine) Run() {
	e.server.run()

	quit := make(chan os.Signal)
	// SIGINT 程序终止(interrupt)信号, 在用户键入INTR字符(通常是Ctrl+C)时发出，用于通知前台进程组终止进程
	// SIGTERM 礼貌地要求进程终止.它将正常终止,清理所有资源(文件,套接字,子进程等),删除临时文件等
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	sig := <-quit
	fmt.Printf("接收到信号[%s]", sig.String())
	//logger.WithContext(ctx).Infof("接收到信号[%s]", sig.String())
	//
	//logger.WithContext(ctx).Infof("服务退出")
}

func (e *engine) ServerHandler() http.Handler {
	return e.server.handler
}




