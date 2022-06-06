// Author: yann
// Date: 2022/5/11
// Desc: ginwrap

package ginwrap

import (
	"github.com/gin-gonic/gin"
)

type GroupWrap struct {
	*gin.RouterGroup
}

func NewGroupWrap(routerGroup *gin.RouterGroup) *GroupWrap {
	return &GroupWrap{RouterGroup: routerGroup}
}

func (g *GroupWrap) Group(relativePath string) *GroupWrap {
	return NewGroupWrap(g.RouterGroup.Group(relativePath))
}

func (g *GroupWrap) Handle(httpMethod, relativePath string, base IBase, handle Handle) {
	g.RouterGroup.Handle(httpMethod, relativePath, func(ctx *gin.Context) {
		var next IBase
		if base != nil {
			var err *DecodeErr
			if next, err = base.DecodeType().Decode(ctx, base); err != nil {
				err.Response().ReturnResult(ctx)
				return
			}
		}
		handle(ctx, next).ReturnResult(ctx)
	})
}

func (g *GroupWrap) Any(relativePath string, base IBase, handle Handle) {
	g.RouterGroup.Any(relativePath, func(ctx *gin.Context) {
		var next IBase
		if base != nil {
			var err *DecodeErr
			if next, err = base.DecodeType().Decode(ctx, base); err != nil {
				err.Response().ReturnResult(ctx)
				return
			}
		}
		handle(ctx, next).ReturnResult(ctx)
	})
}

func (g *GroupWrap) GET(relativePath string, base IBase, handle Handle) {
	g.RouterGroup.GET(relativePath, func(ctx *gin.Context) {
		var next IBase
		if base != nil {
			var err *DecodeErr
			if next, err = base.DecodeType().Decode(ctx, base); err != nil {
				err.Response().ReturnResult(ctx)
				return
			}
		}
		handle(ctx, next).ReturnResult(ctx)
	})
}

func (g *GroupWrap) POST(relativePath string, base IBase, handle Handle) {
	g.RouterGroup.POST(relativePath, func(ctx *gin.Context) {
		var next IBase
		if base != nil {
			var err *DecodeErr
			if next, err = base.DecodeType().Decode(ctx, base); err != nil {
				err.Response().ReturnResult(ctx)
				return
			}
		}
		handle(ctx, next).ReturnResult(ctx)
	})
}

func (g *GroupWrap) DELETE(relativePath string, base IBase, handle Handle) {
	g.RouterGroup.DELETE(relativePath, func(ctx *gin.Context) {
		var next IBase
		if base != nil {
			var err *DecodeErr
			if next, err = base.DecodeType().Decode(ctx, base); err != nil {
				err.Response().ReturnResult(ctx)
				return
			}
		}
		handle(ctx, next).ReturnResult(ctx)
	})
}

func (g *GroupWrap) PATCH(relativePath string, base IBase, handle Handle) {
	g.RouterGroup.PATCH(relativePath, func(ctx *gin.Context) {
		var next IBase
		if base != nil {
			var err *DecodeErr
			if next, err = base.DecodeType().Decode(ctx, base); err != nil {
				err.Response().ReturnResult(ctx)
				return
			}
		}
		handle(ctx, next).ReturnResult(ctx)
	})
}

func (g *GroupWrap) PUT(relativePath string, base IBase, handle Handle) {
	g.RouterGroup.PUT(relativePath, func(ctx *gin.Context) {
		var next IBase
		if base != nil {
			var err *DecodeErr
			if next, err = base.DecodeType().Decode(ctx, base); err != nil {
				err.Response().ReturnResult(ctx)
				return
			}
		}
		handle(ctx, next).ReturnResult(ctx)
	})
}

func (g *GroupWrap) OPTIONS(relativePath string, base IBase, handle Handle) {
	g.RouterGroup.OPTIONS(relativePath, func(ctx *gin.Context) {
		var next IBase
		if base != nil {
			var err *DecodeErr
			if next, err = base.DecodeType().Decode(ctx, base); err != nil {
				err.Response().ReturnResult(ctx)
				return
			}
		}
		handle(ctx, next).ReturnResult(ctx)
	})
}

func (g *GroupWrap) HEAD(relativePath string, base IBase, handle Handle) {
	g.RouterGroup.HEAD(relativePath, func(ctx *gin.Context) {
		var next IBase
		if base != nil {
			var err *DecodeErr
			if next, err = base.DecodeType().Decode(ctx, base); err != nil {
				err.Response().ReturnResult(ctx)
				return
			}
		}
		handle(ctx, next).ReturnResult(ctx)
	})
}
