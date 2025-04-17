package model

type HttpRequest struct {
	Method  string
	Url     string
	Headers map[string][]string
	Body    []byte
}
