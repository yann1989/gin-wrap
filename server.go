// Author: yann
// Date: 2022/5/20
// Desc: ginwrap

package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	ginwrap "github.com/yann1989/gin-wrap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type User struct {
	Name string `json:"name" form:"name" uri:"name" binding:"required" err:"用户名称不能为空"`
}

func (u User) DecodeType() ginwrap.DecodeType {
	return ginwrap.DecodeTypeQuery
}

func main() {
	engine := gin.New()
	wrap := ginwrap.NewEngineWrap(engine, ginwrap.EngineWrapLoggerOption, ginwrap.PrintReqParamsOption, ginwrap.PrintRespParamsOption)
	wrap.GET("", &User{}, func(ctx *gin.Context, base ginwrap.IBase) *ginwrap.Response {
		time.Sleep(time.Second)
		var user = base.(*User)
		return ginwrap.Response200(user.Name)
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
