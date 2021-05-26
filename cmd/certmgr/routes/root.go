package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tangx/srv-lego-certmgr/cmd/certmgr/global"
	"github.com/tangx/srv-lego-certmgr/cmd/certmgr/routes/cert"
	"github.com/tangx/srv-lego-certmgr/cmd/certmgr/routes/certgen"
	"github.com/tangx/srv-lego-certmgr/cmd/certmgr/routes/healthy"
	"github.com/tangx/srv-lego-certmgr/static"
)

var root *gin.RouterGroup

func AppendRoute(e *gin.Engine) {

	// 静态页面
	e.StaticFS("./index", http.FS(static.Index))

	root = e.Group(global.Appname)

	for provider := range global.Providers {
		certgen.AppendRoute(root, provider)
	}

	cert.AppendRoute(root)
	healthy.AppendRoute(root)
}
