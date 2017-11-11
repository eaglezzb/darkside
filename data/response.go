package data


type Response map[string]interface{}

func NewResponse() Response {
	return make(Response)
}

func ErrorResponse(errno int, msg string) Response {
	r := make(Response)
	r.SetErrorInfo(errno, msg)
	return r
}

func (s Response) SetErrorInfo(errno int, msg string) {
	s["errno"] = errno
	s["msg"] = msg
}

func (s Response) AddResponseInfo(key string, val interface{}) {
	s[key] = val
}
