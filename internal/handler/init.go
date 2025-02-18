package handler

import "simple-web-go/internal/service"

type Handlers struct {
	SysUserHandler *SysUserHandler
	// 可以添加其他handler
	SysMenuHandler *SysMenuHandler
	SysRoleHandler *SysRoleHandler
}

func InitHandlers(services *service.Services) *Handlers {
	return &Handlers{
		SysUserHandler: NewSysUserHandler(services.SysUserService),
		// 初始化其他handler
		SysMenuHandler: NewSysMenuHandler(services.SysMenuService),
		SysRoleHandler: NewSysRoleHandler(services.SysRoleService),
	}
}
