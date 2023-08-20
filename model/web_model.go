package model

type WebRequest struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	Author string `json:"author"`
}

type WebResponse struct {
	Code    int32
	Message string
	Data    interface{}
}
