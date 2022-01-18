package router

import (
	"net/http"

	"github.com/GopherReady/GinRestApi/handler/sd"
	"github.com/GopherReady/GinRestApi/router/middleware"
	"github.com/gin-gonic/gin"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// 	middleware
	g.Use(gin.Recovery())
	// 强制浏览器不使用缓存
	g.Use(middleware.NoCache)
	// 浏览器跨域 OPTIONS 请求设置
	g.Use(middleware.Options)
	// 一些安全设置
	g.Use(middleware.Secure)
	g.Use(mw...)
	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	// The health check handlers
	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	return g
}
