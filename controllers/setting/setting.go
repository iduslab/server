package setting

import (
	"github.com/gin-gonic/gin"
	"github.com/iduslab/backend/db"
	"github.com/iduslab/backend/models/req"
	"github.com/iduslab/backend/utils/res"
)

func GetAllValue(c *gin.Context) {
	r := res.New(c)
	data, err := db.GetAllSetting()
	if err != nil {
		r.SendError(res.ERR_SERVER, "")
		return
	}
	r.Response(res.R{
		"data": data,
	})
}

func GetValue(c *gin.Context) {
	r := res.New(c)
	name := c.Param("name")
	data, err := db.GetSetting(name)
	if err != nil {
		r.SendError(res.ERR_SERVER, "")
		return
	}
	r.Response(res.R{
		"data": data,
	})
}

func UpdateValue(c *gin.Context) {
	r := res.New(c)
	name := c.Param("name")
	body := c.MustGet("body").(*req.SettingUpdateValue)
	err := db.UpdateSettingValue(name, body.Value)
	if err != nil {
		r.SendError(res.ERR_SERVER, "")
		return
	}
	r.Response(res.R{})
}

func Add(c *gin.Context) {
	r := res.New(c)
	body := c.MustGet("body").(*req.SettingAddValue)
	err := db.AddSetting(body.Name, body.Description, body.Value)
	if err != nil {
		r.SendError(res.ERR_SERVER, "")
		return
	}
	r.Response(res.R{})
}
