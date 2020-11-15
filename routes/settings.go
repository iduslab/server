package routes

import (
	"github.com/gin-gonic/gin"
	c "github.com/iduslab/backend/controllers/setting"
	m "github.com/iduslab/backend/middlewares"
	"github.com/iduslab/backend/models/req"
)

func setSettingRoutes(r *gin.RouterGroup) {
	r.GET("/", c.GetAllValue)
	r.GET("/:name", c.GetValue)
	r.POST("/", m.VerifyRequest(&req.SettingAddValue{}), c.Add)
	r.PATCH("/:name", m.VerifyRequest(&req.SettingUpdateValue{}), c.UpdateValue)
}
