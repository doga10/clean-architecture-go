package adapters

import (
	"github.com/doga10/clean-architecture-go/src/presentation/protocols"
	"github.com/gin-gonic/gin"
)

func AdaptRouter(controller protocols.Controller) gin.HandlerFunc {
	return func (c *gin.Context) {
		var req protocols.Request
		req.Body = c.Request.Body
		req.Params = c.Params
		req.Headers = c.Request.Header

		res := controller.Handle(req)
		c.JSON(res.Code, res.Data)
	}
}
