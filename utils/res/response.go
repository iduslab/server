package res

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type res struct {
	c *gin.Context
}

type R map[string]interface{}

func New(c *gin.Context) *res {
	return &res{c}
}

func (r *res) Response(data interface{}) {
	m := make(map[string]interface{})
	m["code"] = "SUCCESS"
	m["message"] = ""
	j, _ := json.Marshal(data)
	json.Unmarshal(j, &m)
	r.c.JSON(http.StatusOK, m)
}
