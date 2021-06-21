package tools

type ResponseOk struct {
	Code int         `json:"id"`
	Data interface{} `json:"code"`
	Msg  string      `json:"msg"`
}

func Ok(data interface{}, msg ...string) ResponseOk {
	return ResponseOk{Code: 200, Data: data, Msg: msg[0]}
}
