package protocols

type Request struct {
	Body    []byte      `json:"body"`
	Headers interface{} `json:"headers"`
	Params  interface{} `json:"params"`
	Query   interface{} `json:"query"`
}

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}
