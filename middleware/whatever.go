package middleware

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

func InsideMiddleware(c *gin.Context) {
	param := c.Param("code")
	code, _ := strconv.Atoi(param)
	if code < 200 {
		c.Status(400)
		c.Error(errors.New("code less than 200"))
		return
	}
	c.JSON(200, gin.H{"code": code})
}
