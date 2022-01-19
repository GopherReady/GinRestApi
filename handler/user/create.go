package user

import (
	"fmt"

	"github.com/GopherReady/GinRestApi/handler"
	"github.com/GopherReady/GinRestApi/model"
	"github.com/GopherReady/GinRestApi/pkg/errno"
	"github.com/GopherReady/GinRestApi/util"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log/lager"
)

// Create creates a new user account.
func Create(c *gin.Context) {
	fmt.Println("User Create function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	u := model.UserModel{
		Username: r.Username,
		Password: r.Password,
	}

	// Validate the data.
	if err := u.Validate(); err != nil {
		handler.SendResponse(c, errno.ErrValidation, nil)
		return
	}

	// Encrypt the user password.
	if err := u.Encrypt(); err != nil {
		handler.SendResponse(c, errno.ErrEncrypt, nil)
		return
	}
	// Insert the user to the database.
	if err := u.Create(); err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	rsp := CreateResponse{
		Username: r.Username,
	}

	// Show the user information.
	handler.SendResponse(c, nil, rsp)
}
