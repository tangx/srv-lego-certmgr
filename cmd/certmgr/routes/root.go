package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tangx/srv-lego-certmgr/cmd/certmgr/global"
	"github.com/tangx/srv-lego-certmgr/cmd/certmgr/routes/cert"
	"github.com/tangx/srv-lego-certmgr/cmd/certmgr/routes/certprovider"
)

var root *gin.RouterGroup

func AppendRoute(e *gin.Engine) {
	root = e.Group("certmgr")

	for provider := range global.Providers {
		certprovider.AppendRoute(root, provider)
	}

	cert.AppendRoute(root)

}
