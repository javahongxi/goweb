package result

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Success(data interface{}) Result {
	return Result{Data: data, Msg: "OK"}
}

func Error(code int, msg string) Result {
	return Result{Code: code, Msg: msg}
}
