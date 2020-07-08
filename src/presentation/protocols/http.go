package protocols

import "io"

type Request struct {
	Body    io.ReadCloser `json:"body"`
	Headers interface{}   `json:"headers"`
	Params  interface{}   `json:"params"`
	Query   interface{}   `json:"query"`
}

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}
