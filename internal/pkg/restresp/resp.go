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

func DefaultResp() *Resp {
	return &Resp{
		Code: CodeSuccess,
		Msg:  "",
		Data: nil,
		RespMeta: RespMeta{
			TraceID: "",
			Caller:  "",
			API:     "",
			Time:    time.Now().UnixMilli(),
		},
	}
}

func DefaultSuccessResp() *Resp {
	resp := DefaultResp()
	resp.Msg = "success"
	return resp
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

func DefaultErrWithMsgResp(msg string) *Resp {
	resp := DefaultErrResp()
	resp.Msg = msg
	return resp
}
