package user

import (
	"strconv"

	"github.com/GopherReady/GinRestApi/handler"
	"github.com/GopherReady/GinRestApi/model"
	"github.com/GopherReady/GinRestApi/pkg/errno"
	"github.com/gin-gonic/gin"
)

// Delete delete an user by the user identifier.
func Delete(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))
	if err := model.DeleteUser(uint64(userId)); err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	handler.SendResponse(c, nil, nil)
}
