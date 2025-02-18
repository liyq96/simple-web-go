package dto

type SysUserDto struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Account  string `json:"account"`
	Password string `json:"password"`
}

type PageSysUserDto struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}
