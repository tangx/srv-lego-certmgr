package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tangx/srv-lego-certmgr/cmd/certmgr/routes/cert"
	qcloud "github.com/tangx/srv-lego-certmgr/cmd/certmgr/routes/dnspod"
)

var root *gin.RouterGroup

func AppendRoute(e *gin.Engine) {
	root = e.Group("certmgr")

	qcloud.AppendRoute(root)
	cert.AppendRoute(root)

}
