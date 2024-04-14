package config

type QiNiu struct {
	Enable    bool   `yaml:"enable" json:"enable"` // 是否启用七牛云存储，不启用则图片保存在本地
	AccessKey string `json:"access_key" yaml:"access_key"`
	SecretKey string `json:"secret_key" yaml:"secret_key"`
	Bucket    string `json:"bucket" yaml:"bucket"` // 存储桶
	CDN       string `json:"cdn" yaml:"cdn"`       // Content Delivery Network，通过绑定自定义静态域名来作为静态资源文件的前缀
	Zone      string `json:"zone" yaml:"zone"`     // 存储的地区
	Size      int64  `json:"size" yaml:"size"`     // 存储图片的大小限制(单位是5MB)
}
