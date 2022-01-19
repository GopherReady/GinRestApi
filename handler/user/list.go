package user

import (
	"github.com/GopherReady/GinRestApi/handler"
	"github.com/GopherReady/GinRestApi/pkg/errno"
	"github.com/GopherReady/GinRestApi/service"
	"github.com/gin-gonic/gin"
)

// List the users in the database.
func List(c *gin.Context) {
	var r ListRequest
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	infos, count, err := service.ListUser(r.Username, r.Offset, r.Limit)
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}

	handler.SendResponse(c, nil, ListResponse{
		TotalCount: count,
		UserList:   infos,
	})
}
