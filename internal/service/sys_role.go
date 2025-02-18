package service

import (
	"simple-web-go/internal/dto"
	"simple-web-go/internal/model"
	"simple-web-go/internal/vo"
	"simple-web-go/pkg/utils/paginate"

	"gorm.io/gorm"
)

type SysRoleService struct {
	db *gorm.DB
}

func NewSysRoleService(db *gorm.DB) *SysRoleService {
	return &SysRoleService{db: db}
}

func (s *SysRoleService) SaveOrUpdateRole(req *dto.SysRoleDto) error {
	role := &model.SysRole{
		Id:     uint16(req.Id),
		Name:   req.Name,
		Status: req.Status,
	}
	error := s.db.Save(role).Error
	if error != nil {
		return error
	}
	return nil
}

func (s *SysRoleService) PageRole(req *dto.PageSysRoleDto) (*paginate.Pagination, error) {
	var roles []model.SysRole
	pagination, err := paginate.Paginate(s.db, req.Page, req.PageSize, &roles)
	if err != nil {
		return nil, err
	}
	resp := make([]vo.PageSysRoleVo, 0, len(roles))
	for _, user := range roles {
		resp = append(resp, vo.PageSysRoleVo{
			Id:   uint(user.Id),
			Name: user.Name,
		})
	}
	pagination.Data = resp
	return pagination, nil
}
