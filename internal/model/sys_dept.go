package model

type SysDept struct {
	Id   uint16 `gorm:"primaryKey" json:"id"`
	Name string `gorm:"size:255;not null" json:"name"`
}

func (SysDept) TableName() string {
	return "sys_dept"
}
