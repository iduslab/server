package routes

import (
	"github.com/gin-gonic/gin"
	c "github.com/iduslab/backend/controllers/util"
	m "github.com/iduslab/backend/middlewares"
)

func setUtilRoutes(r *gin.RouterGroup) {
	r.POST("/sendnotice", m.CheckAuth(), m.IsAdmin(), c.SendNotice)
}
