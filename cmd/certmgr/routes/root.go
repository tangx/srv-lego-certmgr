package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tangx/srv-lego-certmgr/cmd/certmgr/routes/cert"
	"github.com/tangx/srv-lego-certmgr/cmd/certmgr/routes/qcloud"
)

var root *gin.RouterGroup

func AppendRoute(e *gin.Engine) {
	root = e.Group("certmgr")

	qcloud.AppendRoute(root)
	cert.AppendRoute(root)

}
