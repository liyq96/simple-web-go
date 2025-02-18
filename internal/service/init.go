package service

import "gorm.io/gorm"

type Services struct {
	SysUserService *SysUserService
	// 可以添加其他服务
	SysMenuService *SysMenuService
	SysRoleService *SysRoleService
}

func InitServices(db *gorm.DB) *Services {
	return &Services{
		SysUserService: NewSysUserService(db),
		// 初始化其他服务
		SysMenuService: NewSysMenuService(db),
		SysRoleService: NewSysRoleService(db),
	}
}
