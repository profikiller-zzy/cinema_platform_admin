package utils

import (
	"github.com/go-playground/validator/v10"
	"reflect"
)

// GetMsgLabel 对参数进行校验，将产生error的字段中标签为`msg`的内容返回
func GetMsgLabel(err error, obj interface{}) string {
	// 使用时，应该传入结构体的指针
	objType := reflect.TypeOf(obj)
	// 检查提供的错误是否是`validator.ValidationErrors`类型
	if errs, ok := err.(validator.ValidationErrors); ok {
		// 如果是
		for _, e := range errs {
			// 循环每一个报错信息
			// 根据报错字段名，获取结构体的具体字段
			field := e.Field()
			if f, exits := objType.Elem().FieldByName(field); exits {
				// 将通过 FieldByName 方法找到的结构体字段（f）的标签（Tag）中名为 "msg" 的信息赋值给变量 msg
				msg := f.Tag.Get("msg")
				return msg
			}
		}
	}
	return err.Error()
}
