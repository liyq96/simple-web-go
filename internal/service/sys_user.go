package service

import (
	"simple-web-go/internal/dto"
	"simple-web-go/internal/model"
	"simple-web-go/internal/vo"
	"simple-web-go/pkg/utils/paginate"

	"gorm.io/gorm"
)

type SysUserService struct {
	db *gorm.DB
}

func NewSysUserService(db *gorm.DB) *SysUserService {
	return &SysUserService{db: db}
}

func (s *SysUserService) AddOrUpdateUser(req *dto.SysUserDto) error {
	user := &model.SysUser{
		Id:       uint16(req.Id),
		Name:     req.Name,
		Account:  req.Account,
		Phone:    req.Phone,
		Password: req.Password,
	}
	error := s.db.Save(user).Error
	if error != nil {
		return error
	}
	return nil
}

func (s *SysUserService) PageUser(req *dto.PageSysUserDto) (*paginate.Pagination, error) {
	var users []model.SysUser
	pagination, err := paginate.Paginate(s.db, req.Page, req.PageSize, &users)
	if err != nil {
		return nil, err
	}
	resps := make([]vo.PageSysUserVo, 0, len(users))
	// for
	for _, user := range users {
		resps = append(resps, vo.PageSysUserVo{
			Id:      uint(user.Id),
			Name:    user.Name,
			Account: user.Account,
			Phone:   user.Phone,
		})
	}
	pagination.Data = resps
	return pagination, nil
}

func (s *SysUserService) DetailUser(id uint64) (*vo.SysUserVo, error) {
	var user model.SysUser
	err := s.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &vo.SysUserVo{
		Id:      uint(user.Id),
		Name:    user.Name,
		Account: user.Account,
		Phone:   user.Phone,
	}, nil
}
