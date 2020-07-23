package response

import (
	"note/utils/constant"

	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
	Meta interface{} `json:"_meta"`
}

func (g *Gin) Response(httpCode int, errorCode string, data interface{}, meta interface{}) {
	g.C.JSON(httpCode, Response{
		Code: httpCode,
		Msg:  constant.GetMsg(errorCode),
		Data: data,
		Meta: meta,
	})
	return
}
