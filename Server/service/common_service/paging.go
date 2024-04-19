package common_service

import (
	"AfterEnd/global"
	"AfterEnd/model"
	"fmt"
	"gorm.io/gorm"
)

type Option struct {
	model.PageInfo
	Debug   bool
	Likes   []string // 模糊匹配的字段
	Preload []string // 预加载的列表
}

type PageInfoDebug struct {
	model.PageInfo
	Debug bool // 是否打印sql语句
}

// PagingList 对不同数据模型的数据项进行分页，返回指定页的所有数据和所有数据项的数量
func PagingList[T any](model T, option Option) (list []T, count int64, err error) {

	DB := global.Db
	if option.Debug {
		DB = global.Db.Session(&gorm.Session{Logger: global.MysqlLog})
	}
	if option.Sort == "" {
		option.Sort = "created_at desc" // 默认按照时间往前排
	}
	DB = DB.Where(model)
	for index, column := range option.Likes {
		if index == 0 {
			DB.Where(fmt.Sprintf("%s like ?", column), fmt.Sprintf("%%%s%%", option.Key))
			continue
		}
		DB.Or(fmt.Sprintf("%s like ?", column), fmt.Sprintf("%%%s%%", option.Key))
	}

	count = DB.Where(model).Find(&list).RowsAffected
	// 这里的query会受上面查询的影响，需要手动复位
	query := DB.Where(model)
	for _, preload := range option.Preload {
		query = query.Preload(preload)
	}
	offset := (option.PageNum - 1) * option.PageSize
	if offset < 0 {
		offset = 0
	}
	err = query.Limit(option.PageSize).Offset(offset).Order(option.Sort).Find(&list).Error

	return list, count, err
}
