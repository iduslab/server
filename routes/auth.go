package routes

import (
	c "github.com/gangjun06/iduslab/controllers/auth"
	m "github.com/gangjun06/iduslab/middlewares"
	"github.com/gangjun06/iduslab/models/req"
	"github.com/gin-gonic/gin"
)

func setAuthRoutes(r *gin.RouterGroup) {
	r.GET("/link", m.VerifyQuery(&req.AuthURL{}), c.AuthURL)
	r.GET("/auth", m.VerifyQuery(&req.Auth{}), c.Auth)
}
