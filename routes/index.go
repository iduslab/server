package routes

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	setAuthRoutes(r.Group("auth"))
	setSettingRoutes(r.Group("setting"))
}
