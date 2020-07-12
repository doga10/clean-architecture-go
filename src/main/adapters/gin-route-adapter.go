package adapters

import (
	"github.com/doga10/clean-architecture-go/src/presentation/protocols"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func AdaptRouter(controller protocols.Controller) gin.HandlerFunc {
	return func (c *gin.Context) {
		var req protocols.Request
		body, _ := ioutil.ReadAll(c.Request.Body)
		req.Body = body
		req.Params = c
		req.Headers = c.Request.Header

		res := controller.Handle(req)
		c.JSON(res.Code, res.Data)
	}
}
