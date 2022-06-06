// Author: yann
// Date: 2022/5/11
// Desc: middleware

package ginwrap

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"sync/atomic"
	"time"
)

var index uint64

func Print(c *gin.Context) {
	current := atomic.AddUint64(&index, 1)
	begin := time.Now()
	c.Next()
	end := time.Now()
	var f func(string) string
	if c.GetInt(CtxCode) == CodeSuccess {
		f = InfoColorMsg
	} else {
		f = ErrorColorMsg
	}
	fmt.Printf("[请求%d返回] %s | %d | %v | %v | %v | %s \n", current, end.Format(TimeFormatSecond), c.Writer.Status(), end.Sub(begin), f(cast.ToString(c.GetInt(CtxCode))), f(c.GetString(CtxMsg)), c.Request.URL.Path)
}

func InfoColorMsg(msg string) string {
	return fmt.Sprintf(" %c[%d;%d;%dm%s%c[0m ", 0x1B, 1, 0, 32, msg, 0x1B)
}

func ErrorColorMsg(msg string) string {
	return fmt.Sprintf(" %c[%d;%d;%dm%s%c[0m ", 0x1B, 1, 0, 31, msg, 0x1B)
}
