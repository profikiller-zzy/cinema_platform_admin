package config

import (
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

type Config struct {
	Mysql  Mysql  `yaml:"mysql"`  // 数据库配置信息
	System System `yaml:"system"` // 系统配置信息
	Logger Logger `yaml:"logger"` // 日志配置信息
	Jwt    Jwt    `yaml:"jwt"`    // jwt配置信息
	Redis  Redis  `yaml:"redis"`  // redis配置信息
	QiNiu  QiNiu  `yaml:"qi_niu"` // 七牛云配置信息
}

// GetUpToken 获取上传七牛云的简单上传凭证
func (q QiNiu) GetUpToken() string {
	if q.AccessKey == "" || q.SecretKey == "" {
		return ""
	}
	bucket := q.Bucket
	putPolicy := storage.PutPolicy{
		Scope:      bucket,
		ReturnBody: `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket)","name":"$(x:name)"}`,
	}
	// 设置2小时为有效期
	putPolicy.Expires = 7200
	mac := qbox.NewMac(q.AccessKey, q.SecretKey)
	upToken := putPolicy.UploadToken(mac)
	return upToken
}

// GetCfg 获取七牛云服务器文件上传的配置
func (q QiNiu) GetCfg() storage.Config {
	Cfg := storage.Config{}
	// 空间对应的机房：华东-浙江2
	region, _ := storage.GetRegionByID(storage.RegionID(q.Zone))
	Cfg.Region = &region
	// 是否使用https域名
	Cfg.UseHTTPS = false
	// 上传是否使用CDN加速
	Cfg.UseCdnDomains = false

	return Cfg
}
