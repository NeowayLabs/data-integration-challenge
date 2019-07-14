package middleware

import (
	"github.com/coscms/xorm"
	"github.com/gin-gonic/gin"
)

//XormMiddleware middleware
func XormMiddleware(xorm *xorm.Engine) gin.HandlerFunc {

	return func(c *gin.Context) {

		if c.Request.Method != "OPTIONS" {
			c.Set("xorm", xorm)
			c.Next()
		}
	}
}
