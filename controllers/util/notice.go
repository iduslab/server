package util

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iduslab/backend/db"
	"github.com/iduslab/backend/models"
	"github.com/iduslab/backend/utils/res"
)

func SendNotice(c *gin.Context) {
	r := res.New(c)

	value, err := db.GetSetting("noticeWebhook")
	if err != nil {
		r.SendError(res.ERR_SERVER, err.Error())
		return
	}

	data := models.DiscordWebhook{
		Content: "asdf",
	}

	dataByte, err := json.Marshal(&data)
	if err != nil {
		r.SendError(res.ERR_SERVER, err.Error())
		return
	}

	reqBody := bytes.NewBuffer(dataByte)

	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodPost, value, reqBody)

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
