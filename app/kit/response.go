//+build !test

package kit

type ResponseOk struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func Ok(data interface{}, msg ...string) ResponseOk {
	if len(msg) > 0 {
		return ResponseOk{Code: 200, Data: data, Msg: msg[0]}
	}
	return ResponseOk{Code: 200, Data: data, Msg: ""}
}
