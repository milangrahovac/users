package users

type Response struct {
	Error *ResponseError `json:"error"`
}

type ResponseError struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}
