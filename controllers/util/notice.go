package util

import (
	"github.com/gin-gonic/gin"
	"github.com/iduslab/backend/utils/res"
)

func SendNotice(c *gin.Context) {
	r := res.New(c)

	r.Response(res.R{})
}
