package router

import (
	"github.com/danta7/mini-gateway/config"
	"github.com/danta7/mini-gateway/internal/core/routing/proxy"
	"github.com/gin-gonic/gin"
)

// Router 定义路由引擎的接口
type Router interface {
	Setup(r gin.IRouter, httpProxy *proxy.HTTPProxy, cfg *config.Config)
}
