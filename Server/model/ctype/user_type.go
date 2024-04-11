package ctype

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type Role int

const (
	PermissionAdmin      Role = 1 // 超级管理员
	PermissionCinemaUser Role = 2 // 影院用户
	PermissionUser       Role = 3 // 普通用户
)

func (s Role) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s Role) String() string {
	var str string
	switch s {
	case PermissionAdmin:
		str = "管理员"
	case PermissionCinemaUser:
		str = "影院用户"
	case PermissionUser:
		str = "普通用户"
	default:
		str = "其他"
	}
	return str
}

// 自定义数据类型必须实现 Scanner 和 Valuer 接口，以便让 GORM 知道如何将该类型接收、保存到数据库

// Scan 实现 sql.Scanner 接口，Scan 将 value 扫描至 r
func (r *Role) Scan(value interface{}) error {
	v, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal value:", value))
	}

	// strings.EqualFold 用于比较两个字符串是否相等，但是不区分大小写，返回值是一个布尔值
	if strings.EqualFold(string(v), "管理员") {
		*r = PermissionAdmin
	} else if strings.EqualFold(string(v), "影院用户") {
		*r = PermissionCinemaUser
	} else if strings.EqualFold(string(v), "普通用户") {
		*r = PermissionUser
	} else { // 默认为普通用户
		*r = PermissionUser
	}

	return nil
}

// Value 实现 driver.Valuer 接口，Value 返回 json value
func (r Role) Value() (driver.Value, error) {
	var str string
	switch r {
	case PermissionAdmin:
		str = "管理员"
	case PermissionCinemaUser:
		str = "影院用户"
	case PermissionUser:
		str = "普通用户"
	default:
		str = "普通用户"
	}
	return str, nil
}
