package rest

type Response struct {
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Success    bool        `json:"success"`
}

func (app Response) GetSuccess(statusCode int, data interface{}) *Response {
	return &Response{
		StatusCode: statusCode,
		Data:       data,
		Message:    "Success",
		Success:    true,
	}
}

func (app Response) GetError(statusCode int, msg string) *Response {
	return &Response{
		StatusCode: statusCode,
		Message:    msg,
		Success:    false,
	}
}

func (app Response) NotAuthorized() *Response {
	return &Response{
		StatusCode: 400,
		Message:    "Not Authorized",
		Success:    false,
	}
}
