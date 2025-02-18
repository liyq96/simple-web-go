package handler

import (
	"net/http"
	"simple-web-go/internal/dto"
	"simple-web-go/internal/service"
	"simple-web-go/pkg/utils/resp"

	"github.com/gin-gonic/gin"
)

type SysRoleHandler struct {
	roleService *service.SysRoleService
}

func NewSysRoleHandler(roleService *service.SysRoleService) *SysRoleHandler {
	return &SysRoleHandler{roleService: roleService}
}

// 增改角色
func (h *SysRoleHandler) SaveOrUpdateRole(c *gin.Context) {
	var req dto.SysRoleDto
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, resp.Error(-1, err.Error()))
		return
	}
	error := h.roleService.SaveOrUpdateRole(&req)
	if error != nil {
		c.JSON(http.StatusBadRequest, resp.Error(-1, error.Error()))
		return
	}
	c.JSON(http.StatusOK, resp.Success(nil))
}

// 分页列表
func (h *SysRoleHandler) PageRole(c *gin.Context) {
	var req dto.PageSysRoleDto
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, resp.Error(-1, err.Error()))
	}
	resps, error := h.roleService.PageRole(&req)
	if error != nil {
		c.JSON(http.StatusInternalServerError, resp.Error(-1, error.Error()))
	}
	c.JSON(http.StatusOK, resp.Success(resps))
}
