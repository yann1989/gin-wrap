// Author: yann
// Date: 2022/5/20
// Desc: ginwrap

package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/HiData-xyz/HiData.More/HiExtern/lib/logging"
	"github.com/gin-gonic/gin"
	ginwrap "github.com/yann1989/gin-wrap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type User struct {
	Name string `json:"name" form:"name" uri:"name"`
}

func (u User) DecodeType() ginwrap.DecodeType {
	return ginwrap.DecodeTypeQuery
}

var log = logging.Logger("LOG-TEST")

func main() {
	engine := gin.New()
	wrap := ginwrap.NewEngineWrap(engine, ginwrap.EngineWrapLoggerOption(), ginwrap.PrintReqParamsOption(), ginwrap.PrintRespParamsOption())
	wrap.GET("", User{}, func(ctx *gin.Context, base ginwrap.IBase) *ginwrap.Response {
		log.Warnw("base:", "args", base)
		return ginwrap.Response200(errors.New("服务异常"))
	})
	server := &http.Server{
		Addr:    ":80",
		Handler: engine,
	}
	go func() {
		quit := make(chan os.Signal)
		signal.Notify(quit, os.Interrupt, os.Kill, syscall.SIGTERM)
		select {
		case s := <-quit:
			fmt.Println("收到停止信号:", s)
		}
		server.Shutdown(context.Background())
	}()

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		panic("启动异常: " + err.Error())
	}
	fmt.Println("正常退出")
}
