package util

type New struct {
	Code int         `json:"code" form:"code"`
	Msg  string      `json:"msg" code:"msg"`
	Data interface{} `json:"data", form:"data"`
}

type Con struct {
	Item  interface{} `json:"item" form:"item"`
	Count int         `json:"count" form:"count"`
}

const (
	Success int = 200
	Fail    int = 300
)

func NewResult(code int, msg string, data ...interface{}) New {
	if len(data) > 0 {
		res := New{
			Code: code,
			Msg:  msg,
			Data: data[0],
		}
		return res
	} else {
		res := New{
			Code: code,
			Msg:  msg,
		}
		return res
	}

}

func NewPage(code int, msg string, data ...interface{}, count int) New {
	res := New{
		Code: code,
		Msg:  msg,
		Data: Con{
			Item:  data,
			Count: count,
		},
	}
	return res
}
