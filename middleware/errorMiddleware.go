package middleware

import (
	"github.com/gin-gonic/gin"
)

func ErrorMiddleware(c *gin.Context) {
	c.Next()
	errs := c.Errors
	if len(errs) > 0 {
		err := errs[0]
		status := c.Writer.Status()
		c.AbortWithStatusJSON(status, gin.H{"error": err.Error()})
	}
}
