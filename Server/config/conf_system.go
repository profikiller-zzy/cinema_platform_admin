package config

import (
	"fmt"
)

type System struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	Env  string `yaml:"env"` // debug为开发者模式，打印所有mysql语句方便调试
}

// Addr 返回服务器地址和端口组成的字符串
func (s System) Addr() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}
