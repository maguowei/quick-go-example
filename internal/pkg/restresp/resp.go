package restresp

import "time"

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

func DefaultSuccessResp() *Resp {
	return &Resp{
		Code: CodeSuccess,
		Msg:  "",
		RespMeta: RespMeta{
			TraceID: "",
			Caller:  "",
			API:     "",
			Time:    time.Now().UnixMilli(),
		},
	}
}

func SuccessWithDataResp(data interface{}) *Resp {
	resp := DefaultSuccessResp()
	resp.Data = data
	return resp
}

func SuccessWithMsgResp(msg string) *Resp {
	resp := DefaultSuccessResp()
	resp.Msg = msg
	return resp
}

func SuccessResp(msg string, data interface{}) *Resp {
	resp := DefaultSuccessResp()
	resp.Msg = msg
	resp.Data = data
	return resp
}

func DefaultErrResp() *Resp {
	return &Resp{
		Code: CodeError,
		Msg:  "",
	}
}
