package routes

import (
	"github.com/gin-gonic/gin"
	c "github.com/iduslab/backend/controllers/util"
	m "github.com/iduslab/backend/middlewares"
	"github.com/iduslab/backend/models/req"
)

func setUtilRoutes(r *gin.RouterGroup) {
	r.POST("/sendnotice", m.VerifyRequest(&req.UtilSendNotice{}), m.CheckAuth(), m.IsAdmin(), c.SendNotice)
}
