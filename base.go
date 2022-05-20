// Author: yann
// Date: 2022/5/11
// Desc: ginwrap

package ginwrap

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"reflect"
)

type IBase interface {
	DecodeType() DecodeType
}

//GetError 自定义错误消息
func GetError(err error, r interface{}) error {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		return err
	}
	s := reflect.TypeOf(r)
	if s.Kind() == reflect.Ptr {
		s = s.Elem()
	}
	if s.Kind() != reflect.Struct {
		return errs
	}
	for _, fieldError := range errs {
		filed, _ := s.FieldByName(fieldError.Field())
		errTag := fieldError.Tag() + "_err"
		// 获取对应binding得错误消息
		errTagText := filed.Tag.Get(errTag)
		// 获取统一错误消息
		errText := filed.Tag.Get("err")
		if errTagText != "" {
			return errors.New(errTagText)
		}
		if errText != "" {
			return errors.New(errText)
		}
		return errors.New(fieldError.Field() + ":" + fieldError.Tag())

	}
	return errors.New("")
}
