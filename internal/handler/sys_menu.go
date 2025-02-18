package handler

import (
	"net/http"

	"simple-web-go/internal/model"
	"simple-web-go/internal/service"

	"github.com/gin-gonic/gin"
)

type SysMenuHandler struct {
	sysMenuService *service.SysMenuService
}

func NewSysMenuHandler(sysMenuService *service.SysMenuService) *SysMenuHandler {
	return &SysMenuHandler{sysMenuService: sysMenuService}
}

func (h *SysMenuHandler) GetMenu(c *gin.Context) {
	var menus []model.SysMenu
	// result := nil
	// if result.Error != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"error": result.Error.Error(),
	// 	})
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{
		"data": menus,
	})
}
