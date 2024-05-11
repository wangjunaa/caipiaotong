package _type

type Response struct {
	success bool
	code    int
	message string
	data    any
}

func SuccessResp(code int, message string, data any) Response {
	resp := Response{
		success: true,
		code:    code,
		message: message,
		data:    data,
	}
	return resp
}
func ErrResp(code int, message string) Response {
	resp := Response{
		success: false,
		code:    code,
		message: message,
	}
	return resp
}
