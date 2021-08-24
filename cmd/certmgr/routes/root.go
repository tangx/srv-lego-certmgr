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

	// 跨域
	e.Use(cros())

	// 静态页面
	e.StaticFS("./index", http.FS(static.Index))
	root = e.Group(global.Appname)

	for provider := range global.Providers {
		certgen.AppendRoute(root, provider)
	}

	cert.AppendRoute(root)
	healthy.AppendRoute(root)
}

// 跨域
func cros() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		if method != "" {
			c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization,X-Token")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
	}
}
