package model

type SysUser struct {
	Id       uint16 `gorm:"primaryKey" json:"id"`
	Name     string `gorm:"size:255;not null" json:"name"`
	Phone    string `gorm:"size:255;not null;" json:"phone"`
	Account  string `gorm:"size:255;not null;" json:"account"`
	Password string `gorm:"size:255;not null" json:"password"`
}

func (SysUser) TableName() string {
	return "sys_user"
}
