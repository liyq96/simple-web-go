package dto

type SysRoleDto struct {
	Id     uint16 `gorm:"primaryKey" json:"id"`
	Name   string `gorm:"size:255;not null" json:"name"`
	Status int    `gorm:"size:255;not null" json:"status"`
}

type PageSysRoleDto struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}
