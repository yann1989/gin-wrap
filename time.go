// Author: yann
// Date: 2022/6/6
// Desc: ginwrap

package ginwrap

import (
	"github.com/gin-gonic/gin"
	"time"
)

func beginTime(c *gin.Context) {
	c.Set(CtxTime, time.Now().UnixNano()/1e6)
	c.Next()
}
