package res

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrType int

const (
	ERR_BAD_REQUEST ErrType = 1 + iota
	ERR_SERVER
	ERR_DUPLICATE
	ERR_AUTH
	ERR_PERMISSION
)

func (r *res) SendError(errType ErrType, text string) {
	var ErrCode string
	var Status int

	set := func(errCode string, status int) {
		ErrCode = errCode
		Status = status
	}

	switch errType {
	case ERR_BAD_REQUEST:
		set("ERR_BAD_REQUEST", http.StatusBadRequest)
	case ERR_SERVER:
		set("ERR_SERVER", http.StatusInternalServerError)
	case ERR_DUPLICATE:
		set("ERR_DUPLICATE", http.StatusConflict)
	case ERR_AUTH:
		set("ERR_AUTH", http.StatusUnauthorized)
	case ERR_PERMISSION:
		set("ERR_PERMISSION", http.StatusForbidden)
	}

	r.c.JSON(Status, gin.H{
		"code":    ErrCode,
		"message": text,
	})
	r.c.Abort()
}
