package restresp

import (
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"time"
)

type Resp struct {
	Code     int         `json:"code"`
	Msg      string      `json:"msg"`
	Data     interface{} `json:"data,omitempty"`
	RespMeta RespMeta    `json:"meta,omitempty"`
}

type RespMeta struct {
	TraceID string `json:"trace_id"`
	Caller  string `json:"caller"`
	API     string `json:"api"`
	Time    int64  `json:"time"`
}

func DefaultResp(c *gin.Context) *Resp {
	return &Resp{
		Code: CodeSuccess,
		Msg:  "",
		Data: nil,
		RespMeta: RespMeta{
			TraceID: requestid.Get(c),
			Caller:  "",
			API:     "",
			Time:    time.Now().UnixMilli(),
		},
	}
}

func DefaultSuccessResp(c *gin.Context) *Resp {
	resp := DefaultResp(c)
	resp.Msg = "success"
	return resp
}

func SuccessWithDataResp(c *gin.Context, data interface{}) *Resp {
	resp := DefaultSuccessResp(c)
	resp.Data = data
	return resp
}

func SuccessWithMsgResp(c *gin.Context, msg string) *Resp {
	resp := DefaultSuccessResp(c)
	resp.Msg = msg
	return resp
}

func SuccessResp(c *gin.Context, msg string, data interface{}) *Resp {
	resp := DefaultSuccessResp(c)
	resp.Msg = msg
	resp.Data = data
	return resp
}

func DefaultErrResp(c *gin.Context) *Resp {
	return &Resp{
		Code: CodeError,
		Msg:  "",
	}
}

func DefaultErrWithMsgResp(c *gin.Context, msg string) *Resp {
	resp := DefaultErrResp(c)
	resp.Msg = msg
	return resp
}
