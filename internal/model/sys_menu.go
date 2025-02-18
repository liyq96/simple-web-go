package model

import "gorm.io/gorm"

type SysMenu struct {
	gorm.Model
	Name     string `gorm:"size:255;not null" json:"name"`
	Path     string `gorm:"size:255;not null" json:"path"`
	ParentID uint   `gorm:"index" json:"parent_id"`
}

func (SysMenu) TableName() string {
	return "sys_menu"
}
