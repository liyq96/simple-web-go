package routes

import (
	"simple-web-go/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterMenuRoutes(r *gin.RouterGroup, sysMenuHandler *handler.SysMenuHandler) {
	menuGroup := r.Group("/menus")
	{
		menuGroup.GET("/getMenu", sysMenuHandler.GetMenu)
	}
}
