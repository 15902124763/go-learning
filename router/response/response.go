package response

type (
	GlobalErrorCode struct {
		ErrorCode int    `json:"errorCode"`
		Error     string `json:"error"`
	}

	Response struct {
		MoreInfo  string      `json:"moreInfo"`
		Data      interface{} `json:"data"`
		ErrorCode int         `json:"errorCode"`
		Error     string      `json:"error"`
	}
)

func OK() *Response {
	res := &Response{
		MoreInfo:  "",
		Data:      true,
		Error:     "",
		ErrorCode: 200,
	}
	return res
}

func Failure() *Response {
	res := &Response{
		MoreInfo:  "",
		Data:      false,
		Error:     "",
		ErrorCode: 200,
	}
	return res
}

func Of(data interface{}) *Response {
	res := &Response{
		MoreInfo:  "",
		Data:      data,
		Error:     "",
		ErrorCode: 200,
	}
	return res
}

func ErrorMsg(msg string) *Response {
	res := &Response{
		MoreInfo:  msg,
		Error:     msg,
		ErrorCode: 500,
	}
	return res
}

func ParticularMsg(msg string) *Response {
	res := &Response{
		MoreInfo:  msg,
		Error:     msg,
		ErrorCode: 501,
	}
	return res
}

func SuccessMsg(msg string) *Response {
	res := &Response{
		MoreInfo:  "",
		Error:     msg,
		ErrorCode: 200,
	}
	return res
}

func Error(g *GlobalErrorCode) *Response {
	if g == nil {
		g = &GlobalErrorCode{
			ErrorCode: 500,
			Error:     "",
		}
	}

	res := &Response{
		MoreInfo:  g.Error,
		Error:     g.Error,
		ErrorCode: g.ErrorCode,
	}
	return res
}

var (
	ParamError       = GlobalErrorCode{ErrorCode: 500, Error: "参数错误"}
	ServerUknowError = GlobalErrorCode{ErrorCode: 500, Error: "服务器发生未知异常"}
)
