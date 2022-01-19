package user

import (
	"github.com/GopherReady/GinRestApi/handler"
	"github.com/GopherReady/GinRestApi/model"
	"github.com/GopherReady/GinRestApi/pkg/errno"
	"github.com/gin-gonic/gin"
)

// Get gets an user by the user identifier.
func Get(c *gin.Context) {
	username := c.Param("username")
	// Get the user by the `username` from the database.
	user, err := model.GetUser(username)
	if err != nil {
		handler.SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	handler.SendResponse(c, nil, user)
}
