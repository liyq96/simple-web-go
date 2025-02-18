package vo

type SysUserVo struct {
	Id      uint   `json:"id"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Account string `json:"account"`
}

type PageSysUserVo struct {
	Id      uint   `json:"id"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Account string `json:"account"`
}
