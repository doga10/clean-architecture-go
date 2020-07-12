package protocols

import "github.com/gin-gonic/gin"

type Request struct {
	Body    []byte       `json:"body"`
	Headers interface{}  `json:"headers"`
	Params  *gin.Context `json:"params"`
	Query   interface{}  `json:"query"`
}

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}
