package model

type SysRoleMenu struct {
	Id     uint16 `gorm:"primaryKey" json:"id"`
	RoleId uint16 `gorm:"size:255;not null" json:"roleId"`
	MenuId uint16 `gorm:"size:255;not null" json:"menuId"`
}

func (SysRoleMenu) TableName() string {
	return "sys_role_menu"
}
