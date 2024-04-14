package utils

import (
	"crypto/md5"
	"fmt"
	"github.com/go-playground/validator/v10"
	"reflect"
)

// GetMsgLabel 对参数进行校验，将产生error的字段中标签为`msg`的内容返回
func GetMsgLabel(err error, obj interface{}) string {
	// 使用时，应该传入结构体的指针
	objType := reflect.TypeOf(obj)
	// 将err接口断言为具体类型
	if errs, ok := err.(validator.ValidationErrors); ok {
		// 断言成功
		for _, e := range errs {
			// 循环每一个报错信息
			// 根据报错字段名，获取结构体的具体字段
			field := e.Field()
			if f, exits := objType.Elem().FieldByName(field); exits {
				msg := f.Tag.Get("msg")
				return msg
			}
		}
	}
	return err.Error()
}

// IsInStringList 判断输入的字符串是否在输入的字符串切片中
func IsInStringList(str string, strList []string) bool {
	for _, suffix := range strList {
		if str == suffix {
			return true
		}
	}
	return false
}

// MD5 将输入源文件进行MD5 hash，返回以16进制表示的128位hash值
func MD5(src []byte) string {
	return fmt.Sprintf("%x", md5.Sum(src))
}
