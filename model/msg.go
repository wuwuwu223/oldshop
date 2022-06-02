package model

type OkMsgWithData struct {
	Msg  string      `json:"msg"`
	Code bool        `json:"code"`
	Data interface{} `json:"data,omitempty"`
}

func NewOkMsgWithData(msg string, code bool, data interface{}) *OkMsgWithData {
	return &OkMsgWithData{
		Msg:  msg,
		Code: code,
		Data: data,
	}
}
