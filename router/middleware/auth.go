package middleware

import (
	"github.com/GopherReady/GinRestApi/handler"
	"github.com/GopherReady/GinRestApi/pkg/errno"
	"github.com/GopherReady/GinRestApi/pkg/token"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse the json web token.
		if _, err := token.ParseRequest(c); err != nil {
			handler.SendResponse(c, errno.ErrTokenInvalid, nil)
			c.Abort()
			return
		}

		c.Next()
	}
}
