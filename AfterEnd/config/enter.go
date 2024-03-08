package config

type Config struct {
	Mysql  Mysql  `yaml:"mysql"`  // 数据库配置信息
	System System `yaml:"system"` // 系统配置信息
	Logger Logger `yaml:"logger"` // 日志配置信息
	Jwt    Jwt    `yaml:"jwt"`    // jwt配置信息
}
