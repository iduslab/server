package util

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/iduslab/backend/models/req"

	"github.com/gin-gonic/gin"
	"github.com/iduslab/backend/db"
	"github.com/iduslab/backend/utils/res"
)

func SendNotice(c *gin.Context) {
	r := res.New(c)

	body := c.MustGet("body").(*req.UtilSendNotice)

	value, err := db.GetSetting("noticeWebhook")
	if err != nil {
		r.SendError(res.ERR_SERVER, err.Error())
		return
	}

	dataByte, err := json.Marshal(&body.Content)
	if err != nil {
		r.SendError(res.ERR_SERVER, err.Error())
		return
	}

	reqBody := bytes.NewBuffer(dataByte)

	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodPost, value, reqBody)
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		r.SendError(res.ERR_SERVER, err.Error())
		return
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		r.SendError(res.ERR_SERVER, err.Error())
		return
	}

	if resp.StatusCode != 200 {
		r.SendError(res.ERR_SERVER, string(respBody))
		return
	}

	r.Response(res.R{})
}
