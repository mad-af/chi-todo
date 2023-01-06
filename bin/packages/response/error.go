package response

func ReplyError(code int, status , message string, reply *Response) {
	(*reply) = Response{Code: code, Status: status, Message: message}
	// return reply
}

func (e *Response) Error() string {
	return e.Message
}


