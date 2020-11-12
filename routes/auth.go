package routes

import (
	"github.com/gin-gonic/gin"
	c "github.com/iduslab/backend/controllers/auth"
	m "github.com/iduslab/backend/middlewares"
	"github.com/iduslab/backend/models/req"
)

func setAuthRoutes(r *gin.RouterGroup) {
	r.GET("/link", m.VerifyQuery(&req.AuthURL{}), c.AuthURL)
	r.GET("/auth", m.VerifyQuery(&req.Auth{}), c.Auth)
}
