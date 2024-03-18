package ctype

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type State int

const (
	Success      State = 1 // 成功
	Fail         State = 2 // 失败
	ToBeReviewed State = 3 // 审核中
)

func (s State) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s State) String() string {
	var str string
	switch s {
	case Success:
		str = "成功"
	case Fail:
		str = "失败"
	case ToBeReviewed:
		str = "审核中"
	default:
		str = "审核中"
	}
	return str
}

// 自定义数据类型必须实现 Scanner 和 Valuer 接口，以便让 GORM 知道如何将该类型接收、保存到数据库

// Scan 实现 sql.Scanner 接口，Scan 将 value 扫描至 r
func (s *State) Scan(value interface{}) error {
	v, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal value:", value))
	}

	// strings.EqualFold 用于比较两个字符串是否相等，但是不区分大小写，返回值是一个布尔值
	if strings.EqualFold(string(v), "成功") {
		*s = Success
	} else if strings.EqualFold(string(v), "失败") {
		*s = Fail
	} else if strings.EqualFold(string(v), "审核中") {
		*s = ToBeReviewed
	} else { // 默认待审核
		*s = ToBeReviewed
	}

	return nil
}

// Value 实现 driver.Valuer 接口，Value 返回 json value
func (s State) Value() (driver.Value, error) {
	var str string
	switch s {
	case Success:
		str = "成功"
	case Fail:
		str = "失败"
	case ToBeReviewed:
		str = "审核中"
	default:
		str = "审核中"
	}
	return str, nil
}
