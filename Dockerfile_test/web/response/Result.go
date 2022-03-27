package response

type Result struct{
	Code int `json:"code"`
	Message string `json:"message"`
	Succeed bool `json:"succeed"`
	Data interface{} `json:"data"`
}

func (r Result) Success(data interface{}) Result{
	return Result{
		Code: 200,
		Message: "成功",
		Succeed: true,
		Data: data,
	}
}

func (r Result) Fail(code int, message string) Result{
	return Result{
		Code: code,
		Message: message,
		Succeed: false,
		Data: nil,
	}
}