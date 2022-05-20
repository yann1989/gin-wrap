// Author: yann
// Date: 2022/5/11
// Desc: ginwrap

package ginwrap

import (
	"github.com/gin-gonic/gin"
)

type engineWrap struct {
	gin.IRouter
}

func NewEngineWrap(IRouter gin.IRouter, opts ...EngineWrapOption) *engineWrap {
	for _, opt := range opts {
		opt()
	}
	IRouter.Use(Print)
	return &engineWrap{IRouter: IRouter}
}

func (e *engineWrap) Group(relativePath string) *GroupWrap {
	return NewGroupWrap(e.IRouter.Group(relativePath))
}

func (e *engineWrap) Handle(httpMethod, relativePath string, base IBase, handle Handle) {
	e.IRouter.Handle(httpMethod, relativePath, func(ctx *gin.Context) {
		if base != nil {
			if err := base.DecodeType().Decode(ctx, base); err != nil {
				err.Response().ReturnResult(ctx)
				return
			}
		}
		handle(ctx, base).ReturnResult(ctx)
	})
	return
}

func (e *engineWrap) Any(relativePath string, base IBase, handle Handle) {
	e.IRouter.Any(relativePath, func(ctx *gin.Context) {
		if base != nil {
			if err := base.DecodeType().Decode(ctx, base); err != nil {
				err.Response().ReturnResult(ctx)
				return
			}
		}
		handle(ctx, base).ReturnResult(ctx)
	})
}

func (e *engineWrap) GET(relativePath string, base IBase, handle Handle) {
	e.IRouter.GET(relativePath, func(ctx *gin.Context) {
		if base != nil {
			if err := base.DecodeType().Decode(ctx, base); err != nil {
				err.Response().ReturnResult(ctx)
				return
			}
		}
		handle(ctx, base).ReturnResult(ctx)
	})
}

func (e *engineWrap) POST(relativePath string, base IBase, handle Handle) {
	e.IRouter.POST(relativePath, func(ctx *gin.Context) {
		if base != nil {
			if err := base.DecodeType().Decode(ctx, base); err != nil {
				err.Response().ReturnResult(ctx)
				return
			}
		}
		handle(ctx, base).ReturnResult(ctx)
	})
}

func (e *engineWrap) DELETE(relativePath string, base IBase, handle Handle) {
	e.IRouter.DELETE(relativePath, func(ctx *gin.Context) {
		if base != nil {
			if err := base.DecodeType().Decode(ctx, base); err != nil {
				err.Response().ReturnResult(ctx)
				return
			}
		}
		handle(ctx, base).ReturnResult(ctx)
	})
}

func (e *engineWrap) PATCH(relativePath string, base IBase, handle Handle) {
	e.IRouter.PATCH(relativePath, func(ctx *gin.Context) {
		if base != nil {
			if err := base.DecodeType().Decode(ctx, base); err != nil {
				err.Response().ReturnResult(ctx)
				return
			}
		}
		handle(ctx, base).ReturnResult(ctx)
	})
}

func (e *engineWrap) PUT(relativePath string, base IBase, handle Handle) {
	e.IRouter.PUT(relativePath, func(ctx *gin.Context) {
		if base != nil {
			if err := base.DecodeType().Decode(ctx, base); err != nil {
				err.Response().ReturnResult(ctx)
				return
			}
		}
		handle(ctx, base).ReturnResult(ctx)
	})
}

func (e *engineWrap) OPTIONS(relativePath string, base IBase, handle Handle) {
	e.IRouter.OPTIONS(relativePath, func(ctx *gin.Context) {
		if base != nil {
			if err := base.DecodeType().Decode(ctx, base); err != nil {
				err.Response().ReturnResult(ctx)
				return
			}
		}
		handle(ctx, base).ReturnResult(ctx)
	})
}

func (e *engineWrap) HEAD(relativePath string, base IBase, handle Handle) {
	e.IRouter.HEAD(relativePath, func(ctx *gin.Context) {
		if base != nil {
			if err := base.DecodeType().Decode(ctx, base); err != nil {
				err.Response().ReturnResult(ctx)
				return
			}
		}
		handle(ctx, base).ReturnResult(ctx)
	})
}
