package paginate

import (
	"gorm.io/gorm"
)

type Pagination struct {
	Page      int         `json:"page"`      // 当前页码
	PageSize  int         `json:"pageSize"`  // 每页记录数
	Total     int64       `json:"total"`     // 总记录数
	TotalPage int         `json:"totalPage"` // 总页数
	Data      interface{} `json:"data"`      // 当前数据
}

// Paginate 分页查询
func Paginate(db *gorm.DB, page, pageSize int, entity interface{}) (*Pagination, error) {
	var total int64
	db.Model(entity).Count(&total) // 查询总记录数

	offset := (page - 1) * pageSize
	err := db.Offset(offset).Limit(pageSize).Find(entity).Error
	if err != nil {
		return nil, err
	}
	totalPage := (int(total) + pageSize - 1) / pageSize // 计算总页数
	return &Pagination{
		Page:      page,
		PageSize:  pageSize,
		Total:     total,
		TotalPage: totalPage,
	}, nil
}
