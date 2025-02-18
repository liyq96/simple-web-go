package handler

import (
	"log"
	"net/http"

	"simple-web-go/internal/dto"
	"simple-web-go/internal/service"
	"simple-web-go/pkg/utils/resp"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SysUserHandler struct {
	userService *service.SysUserService
}

func NewSysUserHandler(userService *service.SysUserService) *SysUserHandler {
	return &SysUserHandler{userService: userService}
}

// 增改用户
func (u *SysUserHandler) SaveOrUpdateUser(c *gin.Context) {
	var req dto.SysUserDto
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, resp.Error(-1, err.Error()))
		return
	}
	error := u.userService.AddOrUpdateUser(&req)
	if error != nil {
		c.JSON(http.StatusBadRequest, resp.Error(-1, error.Error()))
		return
	}
	c.JSON(http.StatusOK, resp.Success(nil))
}

// PageUser 分页获取用户列表
func (u *SysUserHandler) PageUser(c *gin.Context) {
	var req dto.PageSysUserDto
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, resp.Error(-1, err.Error()))
		return
	}
	resps, err := u.userService.PageUser(&req)
	if err != nil {
		log.Printf("处理分页请求失败: %v", err)
		c.JSON(http.StatusInternalServerError, resp.Error(-1, err.Error()))
		return
	}
	c.JSON(http.StatusOK, resp.Success(resps))
}

// 用户详情
func (u *SysUserHandler) DetailUser(c *gin.Context) {
	// var id uint64
	idStr := c.Query("id")
	// 将字符串类型的 id 转换为 uint 类型
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, resp.Error(-1, err.Error()))
		return
	}
	user, err := u.userService.DetailUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, resp.Error(-1, err.Error()))
		return
	}
	c.JSON(http.StatusOK, resp.Success(user))
}
