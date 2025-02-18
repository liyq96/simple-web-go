package service

import (
	"gorm.io/gorm"
)

type SysMenuService struct {
	db *gorm.DB
}

func NewSysMenuService(db *gorm.DB) *SysMenuService {
	return &SysMenuService{db: db}
}
