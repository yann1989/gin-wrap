// Author: yann
// Date: 2022/5/11
// Desc: ginwrap

package ginwrap

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"reflect"
)

type DecodeType uint8

const (
	DecodeTypeJson  DecodeType = iota + 1 //json
	DecodeTypeQuery                       //URL query参数
	DecodeTypeUri                         //URL path参数
	ReqParams       = "请求[%s]: %s\n"
)

type DecodeErr struct {
	Code int
	error
}

func NewDecodeErr(code int, error error) *DecodeErr {
	return &DecodeErr{Code: code, error: error}
}

func (d *DecodeErr) Response() *Response {
	var msg string
	if d.error != nil {
		msg = d.error.Error()
	}
	return &Response{
		Code:    d.Code,
		Message: msg,
	}
}

var (
	ErrNotPoint      = &DecodeErr{CodeServerError, fmt.Errorf("不是指针")}
	ErrDecodeTypeErr = &DecodeErr{CodeServerError, fmt.Errorf("不支持的解码类型")}
	printReq         bool
)

func (d DecodeType) Decode(ctx *gin.Context, dst interface{}) (IBase, *DecodeErr) {
	of := reflect.ValueOf(dst)
	var v reflect.Value
	if of.Kind() != reflect.Ptr {
		v = reflect.New(of.Type())
	} else {
		v = reflect.New(of.Elem().Type())
	}
	var err error
	switch d {
	case DecodeTypeJson:
		err = ctx.ShouldBindJSON(v.Interface())
	case DecodeTypeQuery:
		err = ctx.ShouldBindQuery(v.Interface())
	case DecodeTypeUri:
		err = ctx.ShouldBindUri(v.Interface())
	default:
		return nil, ErrDecodeTypeErr
	}

	if err != nil {
		err = GetError(err, dst)
		return nil, NewDecodeErr(CodeBadRequest, err)
	}

	if of.Kind() != reflect.Ptr {
		v = v.Elem()
	}

	ret := v.Interface().(IBase)

	if printReq {
		indent, _ := json.Marshal(ret)
		fmt.Printf(ReqParams, ctx.Request.URL.Path, string(indent))
	}

	return ret, nil
}
