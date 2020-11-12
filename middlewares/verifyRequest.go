package middlewares

import (
	"github.com/iduslab/backend/utils/res"
	"github.com/gin-gonic/gin"
)

func VerifyRequest(data interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		r := res.New(c)
		if err := c.ShouldBindJSON(data); err != nil {
			r.SendError(res.ERR_BAD_REQUEST, err.Error())
			return
		}
		c.Set("body", data)
	}
}

func VerifyQuery(data interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		r := res.New(c)
		if err := c.ShouldBindQuery(data); err != nil {
			r.SendError(res.ERR_BAD_REQUEST, err.Error())
			return
		}
		c.Set("query", data)
	}
}
