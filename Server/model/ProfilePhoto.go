package model

type ProfilePhoto struct {
	MODEL
	Path string `json:"path"`                // 图片URL，如果存储在本地则为图片路径，存储在云服务器上则是图片链接
	Hash string `json:"hash"`                // 图片的Hash值，用以判断重复图片
	Name string `gorm:"size:36" json:"name"` // 图片的名称
}
