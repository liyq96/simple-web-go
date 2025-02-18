package routes

import (
	"simple-web-go/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.RouterGroup, sysUserHandler *handler.SysUserHandler) {
	userGroup := r.Group("/user")
	{
		userGroup.POST("/saveOrUpdateUser", sysUserHandler.SaveOrUpdateUser)
		userGroup.POST("/pageUser", sysUserHandler.PageUser)
		userGroup.GET("/detailUser", sysUserHandler.DetailUser)
	}
}
