package routes

import (
	"simple-web-go/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterRoleRoutes(r *gin.RouterGroup, sysRoleHandler *handler.SysRoleHandler) {
	userGroup := r.Group("/role")
	{
		userGroup.POST("/saveOrUpdateRole", sysRoleHandler.SaveOrUpdateRole)
		userGroup.POST("/pageRole", sysRoleHandler.PageRole)
		// userGroup.GET("/detailUser", sysRoleHandler.DetailUser)
	}
}
