package service

import (
	"github.com/floriandulzky/http-request-printer/internal/model"
	"io"
	"net/http"
)

type httpServer struct {
	responseChan chan model.HttpRequest
}

func NewHTTPServer(responseChan chan model.HttpRequest) *httpServer {
	return &httpServer{
		responseChan: responseChan,
	}
}

func (p *httpServer) Start() {
	http.ListenAndServe(":8000", p)
}

func (p *httpServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var body []byte
	if r.Body != nil {
		body, _ = io.ReadAll(r.Body)
	}
	p.responseChan <- model.HttpRequest{
		Method:  r.Method,
		Url:     r.URL.String(),
		Headers: r.Header,
		Body:    body,
	}
}
