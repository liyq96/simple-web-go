package model

type SysRole struct {
	Id     uint16 `gorm:"primaryKey" json:"id"`
	Name   string `gorm:"size:255;not null" json:"name"`
	Status int    `gorm:"size:255;not null" json:"status"`
}

func (SysRole) TableName() string {
	return "sys_role"
}
