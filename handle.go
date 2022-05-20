// Author: yann
// Date: 2022/5/11
// Desc: ginwrap

package ginwrap

import (
	"github.com/gin-gonic/gin"
)

type Handle func(ctx *gin.Context, params IBase) *Response
