package auth

import (
	"github.com/danta7/mini-gateway/config"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	authenticator := NewAuthenticator(config.GetConfig())
	return func(c *gin.Context) {
		authenticator.Authenticate(c)
	}
}
